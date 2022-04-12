package main

import (
	"fmt"
)

func main() {
	ChannelDemo7()
}

// 无缓冲通道，必须要有接收才能发送
func ChannelDemo1() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
	// fatal error: all goroutines are asleep - deadlock!
}

// 无缓冲通道，必须有发送才能接收
func ChannelDemo2() {
	ch := make(chan int)
	<-ch
	fmt.Println("接收成功")
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 启动一个 goroutine 来接收值
func recv(c chan int) {
	// recv goroutine 检测到 ch 通道中有值，使用 ret 来接收
	ret := <-c
	fmt.Println("接收成功", ret)
}

func ChannelDemo3() {
	ch := make(chan int)
	go recv(ch)
	// 向 ch 通道中发送值
	ch <- 10
	fmt.Println("发送成功")
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 有缓冲的通道
func ChannelDemo4() {
	ch := make(chan int, 1)
	ch <- 10
	fmt.Printf("发送成功，通道内元素的数量：%v，通道大小：%v\n", len(ch), cap(ch))

	ch <- 20
	//fatal error: all goroutines are asleep - deadlock!
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 单向通道

// out 是一个发送通道，看箭头指向，将值发送到 chan 中
func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

// 发送和接收通道
// in 是一个接收通道，从 chan 中接收值
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

func ChannelDemo5() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 发送通道，循环发送0-9
	go counter(ch1)
	// 接收 ch1 的值，平方处理，使用 ch2 进行发送
	go squarer(ch2, ch1)
	// 接收 ch2 的值
	printer(ch2)
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 向 nil 通道传值永远会阻塞
func ChannelDemo6() {
	var ch chan int
	ch1 := make(chan int)
	fmt.Printf("%#v\n", ch)  // (chan int)(nil)
	fmt.Printf("%#v\n", ch1) //(chan int)(0xc000022120)

	go func() {
		<-ch
	}()
	ch <- 1
	// fatal error: all goroutines are asleep - deadlock!
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func foo(ch chan int) {
	ch <- 1
	ch <- 2
	close(ch)
}

// channel 关闭后依旧可以接收值
func ChannelDemo7() {
	ch := make(chan int)
	go foo(ch)
	// 接收关闭通道的值
	for n := range ch {
		fmt.Println(n)
	}
	fmt.Println("channel now is closed")
}
