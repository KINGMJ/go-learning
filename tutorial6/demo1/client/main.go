package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

func main() {
	//1. 连接远程rpc服务器
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	// 2. 调用方法
	// 求面积
	ret := 0
	err = conn.Call("Rect.Area", Params{5, 10}, &ret)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("面积：", ret)

	err = conn.Call("Rect.Perimeter", Params{5, 10}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("周长：", ret)
}
