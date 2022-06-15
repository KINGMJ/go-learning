package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you're requested: %s\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/**
1. http.ResponseWriter: 这个参数是给你写你的text/html响应的
2. http.Request : 它包含所有HTTP请求的信息，比如URL 或者 header
**/
