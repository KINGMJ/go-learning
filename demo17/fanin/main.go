package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

func main() {
	demo2()
}

func demo1() {
	ch1 := make(chan any)
	ch2 := make(chan any)
	go producer(ch1)
	go producer(ch2)

	resultCh := fanIn(ch1, ch2)
	for v := range resultCh {
		fmt.Println("Received:", v)
	}
}

// 生成一个区间随机数
func randBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

// 生产队列
func producer(ch chan any) {
	for i := 1; i <= 5; i++ {
		ch <- rand.Intn(100)
		// 500ms ~ 2s
		time.Sleep(time.Duration(randBetween(500, 2000)) * time.Millisecond)
	}
	close(ch)
}

// fanIn 模式，从多个通道中接收数据返回到一个通道中
func fanIn(channels ...<-chan any) <-chan any {
	var wg sync.WaitGroup
	out := make(chan any)
	copy := func(c <-chan any) {
		defer wg.Done()
		for v := range c {
			out <- v
		}
	}
	wg.Add(len(channels))
	for _, ch := range channels {
		go copy(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func fanInReflect(chans ...<-chan any) <-chan any {
	out := make(chan any)
	go func() {
		defer close(out)
		// 构造SelectCase slice
		var cases []reflect.SelectCase
		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		// 循环，从cases中选择一个可用的
		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok { // 此channel已经close
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			out <- v.Interface()
		}
	}()
	return out
}

func demo2() {
	ch1 := make(chan any)
	ch2 := make(chan any)
	go producer(ch1)
	go producer(ch2)

	resultCh := fanInReflect(ch1, ch2)
	for v := range resultCh {
		fmt.Println("Received:", v)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
