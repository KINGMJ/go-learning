package main

import (
	"fmt"
	"time"
)

func main() {
	resetTimer()
}

func timerDemo() {
	// 创建一个定时器
	timer := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1: %v\n", t1)

	t2 := <-timer.C
	fmt.Printf("t2: %v\n", t2)
}

/*
	timer 只能响应一次
*/
func timerDemo2() {
	timer := time.NewTimer(1 * time.Second)
	// while true
	for {
		<-timer.C
		fmt.Println("时间到")
	}
	//时间到
	//fatal error: all goroutines are asleep - deadlock!
}

/*
	实现延时的三种方式
*/
func timerDemo3() {
	time.Sleep(time.Second)
	fmt.Println("1s 到")

	timer := time.NewTimer(time.Second)
	<-timer.C
	fmt.Println("2s 到")

	<-time.After(time.Second)
	fmt.Println("3s 到")
}

/*
	停止定时器
*/
func stopTimer() {
	timer := time.NewTimer(time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer 2 fired")
	}()

	stop := timer.Stop()
	if stop {
		fmt.Println("Timer 2 stop")
	}
	time.Sleep(2 * time.Second)
}

func resetTimer() {
	timer := time.NewTimer(3 * time.Second)
	// 重置为 1s
	timer.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer.C)
}
