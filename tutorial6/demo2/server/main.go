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

type Rect struct{}

// RPC 服务方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

// RPC 服务方法，求矩形周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

func main() {
	rectServer := new(Rect)
	// 1. 注册一个rect的服务
	rpc.Register(rectServer)

	// 2. 监听端口
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Panicln(err)
	}

	// 3. 开始建立连接
	for {
		fmt.Println("开始建立连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
