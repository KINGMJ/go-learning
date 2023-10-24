package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Chain(helloHandler, Method("POST"), Logging()))
	http.ListenAndServe(":8080", nil)
}

// Middleware 定义一个中间件类型
//
//	@param http.HandlerFunc
//	@return http.HandlerFunc
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging 日志中间件
//
//	@return Middleware
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

// Method http请求判断中间件
//
//	@param m
//	@return Middleware
func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(
					w,
					fmt.Sprintf("%d %s", http.StatusBadRequest, http.StatusText(http.StatusBadRequest)),
					http.StatusBadRequest,
				)
				return
			}
			f(w, r)
		}
	}
}

// Chain 链接多个中间件
//
//	@param f
//	@param middlewares
//	@return http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
