package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	demo3()
}

func demo1() {
	start := time.Now()
	// 等待 orDone 通道关闭
	<-or(
		sig(1, 10*time.Second),
		sig(2, 20*time.Second),
		sig(3, 30*time.Second),
	)
	fmt.Printf("done after %v", time.Since(start))
}

// or 函数接收多个输入通道，并返回一个输出通道，该输出通道会在任一输入通道关闭时关闭。
func or(channels ...<-chan any) <-chan any {
	// 特殊情况，只有零个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	// 创建 orDone 通道
	orDone := make(chan any)
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

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
func demo2() {
	chs := CreateChannels(1, 10)
	// 创建SelectCase
	var cases = CreateCases(chs...)
	// 执行10次select
	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() { // recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send:", cases[chosen].Dir, ok)
		}
	}
}

// 创建 count 数量的 chan int slice
func CreateChannels(count int, cap int) []chan int {
	channels := make([]chan int, count)
	for i := 0; i < count; i++ {
		channels[i] = make(chan int, cap)
	}
	return channels
}

// createCases 函数分别为每个 chan 生成了 recv case 和 send case，并返回一个 reflect.SelectCase 数组
func CreateCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase
	// 创建 recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	// 创建 send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}
	return cases
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func demo3() {
	start := time.Now()
	// 等待 orDone 通道关闭
	<-or1(
		sig(1, 10*time.Second),
		sig(2, 20*time.Second),
		sig(3, 30*time.Second),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func or1(channels ...<-chan any) <-chan any {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan any)
	go func() {
		defer close(orDone)
		// 利用反射构建 selectCase
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		// 随机选择一个可用的case
		reflect.Select(cases)
	}()
	return orDone
}
