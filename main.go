package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	dir := flag.String("d", ".", "directory to serve from")
	port := flag.String("p", "8080", "port to serve on")
	flag.Parse()

	log := log.New(os.Stdout, "[fileserver] ", log.LstdFlags)

	log.Printf("fileserver running @ :%s\n", *port)
	log.Panic(http.ListenAndServe(":"+*port, loggerMiddleware(log)(http.FileServer(http.Dir(*dir)))))
}

func loggerMiddleware(log *log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			recorder := &statusRecorder{
				ResponseWriter: w,
				statusCode: http.StatusOK,
			}

			before := time.Now()
			next.ServeHTTP(w, r)
			elapsed := time.Since(before)

			log.Printf("%s %s%s | %v | %d \n", r.Method, r.URL.Path, r.URL.RawQuery, elapsed, recorder.statusCode)
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
