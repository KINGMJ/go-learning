package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()

	for {
		var data [1024]byte
		// 接收数据
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read upd failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		// 发送数据
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to upd failed, err:", err)
			continue
		}
	}
}
