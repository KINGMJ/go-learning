package main

import (
	"fmt"
	"time"
)

func asStream(done <-chan struct{}, values ...any) <-chan any {
	//创建一个unbuffered的channel
	s := make(chan any)
	// 启动一个goroutine，往s中塞数据
	go func() {
		// 退出时关闭chan
		defer close(s)
		// 遍历数组
		for _, v := range values {
			select {
			case <-done:
				fmt.Println("通道关闭了...")
				return
			case s <- v: // 将数组元素塞入到chan中
				fmt.Printf("发送数据: %#v 到 s 中\n", v)
			}
		}
	}()
	return s
}

func demo1() {
	data := []any{1, 2, "hello", 4, "jack", 6, true, 8, 9, 10}
	done := make(chan struct{})
	ch := asStream(done, data...)
	for v := range ch {
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
	// 2s 后关闭通道
	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
func takeN(done <-chan struct{}, valueStream <-chan any, num int) <-chan any {
	takeStream := make(chan any)
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}
