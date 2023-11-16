package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

func main() {
	//1. 使用 net.Dial 和 rpc 微服务建立连接
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 2. 建立基于json编码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	// 3. 调用远程函数
	// 求面积
	ret := 0
	err = client.Call("Rect.Area", Params{5, 10}, &ret)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("面积：", ret)

	err = client.Call("Rect.Perimeter", Params{5, 10}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("周长：", ret)
}
