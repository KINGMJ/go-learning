package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var once sync.Once
	var addr = "baidu.com:8080"
	var conn net.Conn
	var err error
	once.Do(func() {
		conn, err = net.Dial("tcp", addr)
		fmt.Println(conn, err)
	})
}
