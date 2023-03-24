package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/KINGMJ/go-learning/tutorial5/demo3/gif"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	if isFavicon(r) {
// 		return
// 	}
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// 	fmt.Printf("handler count：%d\n", count)
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Printf("counter count：%d\n", count)
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
// 	for k, v := range r.Header {
// 		fmt.Fprintf(w, "Header [%q] = %q\n", k, v)
// 	}
// 	fmt.Fprintf(w, "Host = %q\n", r.Host)
// 	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
// 	// 通过 ParseForm 获取 Form
// 	if err := r.ParseForm(); err != nil {
// 		log.Print(err)
// 	}
// 	for k, v := range r.Form {
// 		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
// 	}
// }

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 在 web 服务器显示之前的 gif 图片，并且可以设置 cycles 的值
func handler(w http.ResponseWriter, r *http.Request) {
	if isFavicon(r) {
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	cycles := r.Form.Get("cycles")
	f, err := strconv.ParseFloat(cycles, 64)
	if err != nil {
		log.Fatal(err)
	}
	gif.Lissajous(w, f)
}

// 判断请求是否是 favicon.ico
//
//	@param r
//	@return bool
func isFavicon(r *http.Request) bool {
	return r.URL.RequestURI() == "/favicon.ico"
}
