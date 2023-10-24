package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/foo", loggingMiddleware(fooHandler))
	http.HandleFunc("/bar", loggingMiddleware(barHandler))
	http.ListenAndServe(":8080", nil)
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			fmt.Printf("Request URL: %s, time: %s\n", r.URL.Path, time.Since(start))
		}()
		next(w, r)
	}
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}
