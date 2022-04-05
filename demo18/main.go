package main

import (
	"fmt"
	"time"
)

func main() {
	channelDemo4()
}

// 无缓冲通道，必须要有接收才能发送
func channelDemo1() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
	// fatal error: all goroutines are asleep - deadlock!
}

// 启动一个 goroutine 来接收值
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func channelDemo2() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
	// 接收成功 10
	// 发送成功
}

// 有缓冲的通道
func channelDemo3() {
	ch := make(chan int, 1)
	ch <- 10
	fmt.Printf("发送成功，通道内元素的数量：%v，通道大小：%v\n", len(ch), cap(ch))

	ch <- 20
	//fatal error: all goroutines are asleep - deadlock!
}

// 使用通道做并发同步
func channelDemo4() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")

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

// 循环接收
func channelDemo5() {
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
		close(ch)
	}()

	// 遍历接收通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)
		// 当遇到数据0时, 退出接收循环。如果不退出，会导致 fatal error: all goroutines are asleep - deadlock!
		// if data == 0 {
		// 	break
		// }
	}
}

// 单向通道

// 发送通道
func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

// 发送和接收通道
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

// 接收通道
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func channelDemo6() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 发送通道，循环发送0-9
	go counter(ch1)
	// 接收 ch1 的值，平方处理，使用 ch2 进行发送
	go squarer(ch2, ch1)
	// 接收 ch2 的值
	printer(ch2)
}

func foo(ch chan int) {
	ch <- 1
	ch <- 2
	close(ch)
}

func channelDemo7() {
	ch := make(chan int)
	go foo(ch)
	for n := range ch {
		fmt.Println(n)
	}
	fmt.Println("channel now is closed")
}
