package main

import (
	"fmt"
	"time"
)

// or 函数接收多个输入通道，并返回一个输出通道，该输出通道会在任一输入通道关闭时关闭。
func or(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	// 创建 orDone 通道
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2: // 2个也是一种特殊情况
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: // 超过两个，二分法递归处理
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()
	return orDone
}

func sig(id int, after time.Duration) <-chan any {
	c := make(chan any)
	go func() {
		defer close(c)
		time.Sleep(after)
		fmt.Printf("goroutine #%d done.\n", id)
	}()
	return c
}

func main() {
	start := time.Now()

	orDone := or(
		sig(1, 10*time.Second),
		sig(2, 20*time.Second),
		sig(3, 30*time.Second),
	)

	// 等待 orDone 通道关闭
	<-orDone

	fmt.Printf("done after %v", time.Since(start))
}
