package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	httpServer()
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!\n"))

	// 直接获取是一个空的值，需要ParseForm才能获取到
	fmt.Fprintf(w, "before parse from %v\n", r.Form) // before parse from map[]

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse from error %v\n", r.Form)
	}
	fmt.Fprintf(w, "before parse from %v\n", r.Form) // before parse from map[name:[jack]]
}

func httpServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/**
1. 访问 http://localhost:8080?name=jack
**/
