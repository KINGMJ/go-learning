package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	getDemo()
}

// 写文件
func writeFileDemo() {
	err := os.WriteFile("hello.txt", []byte("Hello, World!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// 读文件
func readFileDemo() {
	filename := "hello.txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	os.Stdout.Write(body)
}

// http server 例子
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是 application/x-www-form-urlencoded 时解析 form 数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是 application/json 时从 r.Body 读取数据
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func httpServerDemo() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", handler)
	http.HandleFunc("/get-with-parameter", getHandler)
	http.HandleFunc("/post", postHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// get 请求
func getDemo() {
	res, err := http.Get("https://www.baidu.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}

// get 请求，带参数
func getWithParameterDemo() {
	apiUrl := "http://localhost:8080/get-with-parameter"
	// url parameters
	// https://pkg.go.dev/net/url#Values
	form := url.Values{}
	form.Set("name", "小王子")
	form.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	// url encode
	u.RawQuery = form.Encode()
	fmt.Println(u.String())
	// http get 请求
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

// post 请求
func postDemo() {
	apiUrl := "http://localhost:8080/post"
	// 使用表单提交
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=小王子&age=18"

	// 使用 json 提交
	contentType := "application/json"
	data := `{"name": "小王子", "age": "18"}`
	resp, err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}
