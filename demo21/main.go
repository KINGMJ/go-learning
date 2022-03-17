package main

import (
	"fmt"
	"time"
)

func main() {
	selectDemo5()
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

// timeout
func selectDemo3() {
	chResult := make(chan int, 1)
	go func() {
		time.Sleep(1 * time.Second)
		chResult <- 5
		fmt.Println("Worker finished")
	}()

	select {
	case res := <-chResult:
		fmt.Println("Got from worker\n", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("Timed out before worker finished\n")
		// Timed out before worker finished
	}
}

// 使用 default 进行非阻塞等待
func selectDemo4() {
	ch := make(chan int, 1)

end:
	for {
		select {
		case n := <-ch:
			fmt.Printf("received %d from a channel\n", n)
			break end
		default:
			fmt.Println("Channel is empty")
			ch <- 8
		}
	}
	//Channel is empty
	//received 8 from a channel
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
func selectDemo5() {
	ch, chQuit := make(chan int), make(chan struct{})
	go worker(ch, chQuit)
	ch <- 3
	chQuit <- struct{}{}
	// wait to be signalled back by the worker
	<-chQuit

	//Got value 3
	//Signalled on quit channel. Finishing
}
