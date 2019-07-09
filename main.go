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
	version, commit, date string
)

func main() {
	dir := flag.String("d", ".", "directory to serve from")
	port := flag.String("p", "8080", "port to serve on")
	versionFlag := flag.Bool("v", false, "fileserver version")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("fileserver version: %s, date: %s, commit: %s\n", version, date, commit)
		os.Exit(0)
	}

	log := log.New(os.Stdout, "[fileserver] ", log.LstdFlags)

	var handler http.Handler = http.FileServer(http.Dir(*dir))
	handler = cacheMiddleware(handler)
	handler = loggerMiddleware(log)(handler)

	log.Printf("fileserver running @ :%s\n", *port)
	log.Panic(http.ListenAndServe(":"+*port, handler))
}

func cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// https://stackoverflow.com/a/2068407
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.

		next.ServeHTTP(w, r)
	})
}

func loggerMiddleware(log *log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			recorder := &statusRecorder{
				ResponseWriter: w,
				statusCode:     http.StatusOK,
			}

			before := time.Now()
			next.ServeHTTP(w, r)
			elapsed := time.Since(before)

			log.Printf("%s %s%s | %v | %d %s \n", r.Method, r.URL.Path, r.URL.RawQuery, elapsed, recorder.statusCode, http.StatusText(recorder.statusCode))
		})
	}
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}
