package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// 监听tcp的 20000 端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		// 建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		// 启动一个Goroutine处理连接
		go process(conn)
	}
}

func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		// 读取数据
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		// 发送数据
		conn.Write([]byte(recvStr))
	}
}
