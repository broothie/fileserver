package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	version, date string
)

func main() {
	dir := flag.String("d", ".", "directory to serve from")
	port := flag.String("p", "8080", "port to serve on")
	versionFlag := flag.Bool("v", false, "show version")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("fileserver v%s; built %s\n", version, date)
		os.Exit(0)
	}

	logger := log.New(os.Stdout, "[fileserver] ", 0)

	handler := http.FileServer(http.Dir(*dir))
	handler = cacheMiddleware(handler)
	handler = loggerMiddleware(logger)(handler)

	logger.Printf("running @ http://localhost:%s\n", *port)
	logger.Panic(http.ListenAndServe(":"+*port, handler))
}

func cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		// https://stackoverflow.com/a/2068407
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
	})
}

func loggerMiddleware(logger *log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			recorder := &recorder{ResponseWriter: w, statusCode: http.StatusOK}

			before := time.Now()
			next.ServeHTTP(w, r)
			elapsed := time.Since(before)

			logger.Printf("%s %s%s %dB | %d %s %dB | %v \n",
				r.Method,
				r.URL.Path,
				r.URL.RawQuery,
				r.ContentLength,
				recorder.statusCode,
				http.StatusText(recorder.statusCode),
				len(recorder.body),
				elapsed,
			)
		})
	}
}

type recorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (r *recorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *recorder) Write(body []byte) (int, error) {
	r.body = body
	return r.ResponseWriter.Write(body)
}
