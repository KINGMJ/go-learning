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
	fmt.Fprintf(w, "before parse from %v\n", r.Form) // before parse from map[]
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse from erro %v\n", r.Form)
	}
	fmt.Fprintf(w, "before parse from %v\n", r.Form) // before parse from map[name:[jack]]
}

func httpServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
