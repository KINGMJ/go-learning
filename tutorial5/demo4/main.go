package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	demo4()
}

func demo1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// resp.Body.Close关闭resp的Body流，防止资源泄露
		defer resp.Body.Close()
		// resp的Body字段包括一个可读的服务器响应流
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			// Exit 函数可以让当前程序以给出的状态码 code 退出。
			// 一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行。
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// 练习 1.7
func demo2() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		// 用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy  failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("The number of bytes are: %d\n", b)
	}
}

// 练习 1.8
func demo3() {
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "http://")
		if !hasPrefix {
			url = "http://" + url
		}
		fmt.Println(url)
	}
}

// 练习 1.9
func demo4() {
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "http://")
		if !hasPrefix {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
	}
}
