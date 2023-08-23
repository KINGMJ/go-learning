package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	demo4()
}

// 使用 gorilla/mux 来管理路由
// 访问：http://localhost:8080/books/go/page/1
// Your're requested the book: go on page 1
func demo1() {
	// 创建一个路由实例
	r := mux.NewRouter()
	// 替代 http 的 HandleFunc 方法
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// 从动态参数中获取数据
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "Your're requested the book: %s on page %s\n", title, page)
	})

	// 可以通过正则进行限制，如果不是下面这种格式都是404
	// 访问：http://localhost:8080/articles/demo/1
	// Category: demo
	// Id: 1
	r.HandleFunc("/articles/{category}/{id:[0-9+]}", ArticleHandler)
	http.ListenAndServe(":8080", r)
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n Id: %v\n", vars["category"], vars["id"])
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 限定请求
func demo2() {
	r := mux.NewRouter()
	// 1. 限定方法：只能是 post 请求，如果用其他的方法访问，会得到一个 405 错误
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	// 2. 限定域名或子域名
	r.HandleFunc("/books/{title}", ListBook).Host("www.goexample.com")
	r.HandleFunc("/books/{title}", ListBook).Host("{subdomain:[a-z]+}.example.com")
	http.ListenAndServe(":8080", r)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, vars["title"])
}

func ListBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, vars["title"])
}

func demo3() {
	r := mux.NewRouter()
	// 子路由限定
	bookRouter := r.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Your're requested all book")
	})
	bookRouter.HandleFunc("/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		fmt.Fprintf(w, "Your're requested the book: %s\n", title)
	})
	http.ListenAndServe(":8080", r)
}

func demo4() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		fmt.Fprintln(w, params)
		query := params.Get("name")
		fmt.Fprintf(w, "Search Query: %s", query)
	})
	http.ListenAndServe(":8080", r)
}
