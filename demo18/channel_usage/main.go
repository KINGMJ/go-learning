package main

import (
	"fmt"
	"time"
)

func main() {
	ChannelDemo2()
}

// 循环接收值
func ChannelDemo1() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		// 从3循环到0
		for i := 3; i >= 0; i-- {
			// 发送3到0之间的数值
			ch <- i
			// 每次发送完时等待1s
			time.Sleep(time.Second)
		}
		// 发送完毕关闭通道，如果不关闭的话，接收方数据接收完了继续接收会导致错误
		// close(ch)
	}()

	// 遍历接收通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)
		// 当遇到数据0时, 退出接收循环。如果不退出，会导致 fatal error: all goroutines are asleep - deadlock!
		if data == 0 {
			break
		}
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 使用通道做并发同步
func ChannelDemo2() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		time.Sleep(2 * time.Second)
		// 通过通道通知 main 的 goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	// 等待匿名goroutine
	<-ch
	fmt.Println("all done")

	/*
		wait goroutine
		start goroutine
		exit goroutine
		all done
	*/
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 信号通道，通过 chan struct{} 发送无意义的值
func worker(ch chan int, chQuit chan struct{}) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("Got value %d\n", v)
		case <-chQuit:
			fmt.Printf("Signalled on quit channel. Finishing\n")
			chQuit <- struct{}{}
			return
		}
	}
}

func ChannelDemo3() {
	ch, chQuit := make(chan int), make(chan struct{})
	go worker(ch, chQuit)
	ch <- 3
	chQuit <- struct{}{}
	// wait to be signalled back by the worker
	<-chQuit

	//Got value 3
	//Signalled on quit channel. Finishing
}
