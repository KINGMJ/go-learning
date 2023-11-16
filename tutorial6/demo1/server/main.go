package main

import (
	"log"
	"net/http"
	"net/rpc"
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
	rect := new(Rect)
	// 1. 注册一个rect的服务
	rpc.Register(rect)
	// 2. 将服务绑定到 http 协议上
	rpc.HandleHTTP()
	// 3. 监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
}
