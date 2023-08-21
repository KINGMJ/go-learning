package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to home page")
}

type HomeHandler struct{}

func (s *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to home2 page")
}

func main() {
	http.HandleFunc("/", rootHandler)
	// 创建一个 HomeHandler 处理器，并将其与 /home 路径关联
	staticHandler := &HomeHandler{}
	http.Handle("/home", staticHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
