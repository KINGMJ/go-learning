package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
)

// 使用 r.FormValue() 获取值
func formValueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.FormValue("name"))
	fmt.Fprintln(w, r.PostFormValue("name"))
}

// r.Form、r.PostForm 等
func formHandler(w http.ResponseWriter, r *http.Request) {
	// 获取上传的文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 打印文件信息
	fmt.Fprintf(w, "File Name: %s\n", handler.Filename)
	fmt.Fprintf(w, "File Size: %d\n", handler.Size)
}

// r.Body
func bodyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 读不到 body 了
	body2, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "read the data: %s \n", body)
	fmt.Fprintf(w, "read the body2: %s", body2)
	fmt.Fprintln(w, reflect.TypeOf(body2))
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func bodyHandler2(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// 解析JSON数据
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Username: %s\n", user.Username)
	fmt.Fprintf(w, "Email: %s\n", user.Email)
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("Failed to make GET request:", err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status:", response.Status)
}

func main() {
	http.HandleFunc("/", formValueHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/body", bodyHandler2)
	http.HandleFunc("/client", clientHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
