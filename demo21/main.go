package main

import (
	"fmt"
	"time"
)

func main() {
	selectDemo2()
}

// select 实现多路复用
func foo1(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "foo1"
}

func foo2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "foo2"
}

func selectDemo() {
	output1 := make(chan string)
	output2 := make(chan string)

	// 开启两个 goroutine
	go foo1(output1)
	go foo2(output2)

	// 用 select 监控
	select {
	case s1 := <-output1:
		fmt.Println("s1 = ", s1)
	case s2 := <-output2:
		fmt.Println("s2 =", s2)
	default:
		fmt.Println("没有接收到数据")
	}
}

// 多个 channel 同时 ready， 随机选择一个执行
func selectDemo2() {
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	go func() {
		int_chan <- 1
	}()

	go func() {
		string_chan <- "hello"
	}()

	select {
	case s1 := <-int_chan:
		fmt.Println("int: ", s1)
	case s2 := <-string_chan:
		fmt.Println("string: ", s2)
	}
}
