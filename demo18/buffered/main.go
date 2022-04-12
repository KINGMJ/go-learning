package main

import (
	"fmt"
	"time"
)

func main() {
	unbuffered()
	buffered()
}

// 无缓冲通道与缓冲通道

// 生产者，对于偶数生产者 10ms 生产一个产品；奇数生产者 1ms 生产一个产品
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			time.Sleep(10 * time.Millisecond)
		} else {
			time.Sleep(1 * time.Millisecond)
		}
		ch <- i
	}
}

// 消费者，偶数消费者 1ms 消费一个产品；奇数消费者 10ms 消费一个产品
func consumer(ch <-chan int) {
	total := 0
	for i := 0; i < 5; i++ {
		if i%2 == 1 {
			time.Sleep(10 * time.Millisecond)
		} else {
			time.Sleep(1 * time.Millisecond)
		}
		total += <-ch
	}
}

// 无缓冲通道
func unbuffered() {
	timeStart := time.Now()
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
	fmt.Printf("Unbuffered version took %s\n", time.Since(timeStart))
}

// 缓冲通道
func buffered() {
	timeStart := time.Now()
	ch := make(chan int, 5)
	go producer(ch)
	consumer(ch)
	fmt.Printf("Buffered version took %s\n", time.Since(timeStart))
}

// Unbuffered version took 56.015488ms
// Buffered version took 35.24638ms

// 可以明显看出缓冲通道效率要高很多
