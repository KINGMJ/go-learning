package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

func main() {
	demo4()
}

func demo1() {
	var once sync.Once
	f1 := func() {
		fmt.Println("in f1")
	}
	f2 := func() {
		fmt.Println("in f2")
	}
	once.Do(f1)
	once.Do(f2)
}

func demo2() {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println(i)
		})
	}
}

func demo3() {
	var once sync.Once
	var addr = "baidu.com:8080"
	var conn net.Conn
	var err error
	once.Do(func() {
		conn, err = net.Dial("tcp", addr)
		fmt.Println(conn, err)
	})
}

func demo4() {
	var once sync.Once
	var googleConn net.Conn // 到Google网站的一个连接

	once.Do(func() {
		// 建立到google.com的连接，有可能因为网络的原因，googleConn并没有建立成功，此时它的值为nil
		googleConn, _ = net.Dial("tcp", "google.com:80")
	})
	// 发送http请求
	googleConn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\n Accept: */*\r\n\r\n"))
	io.Copy(os.Stdout, googleConn)
}
