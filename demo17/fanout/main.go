package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 生成一个区间随机数
func randBetween(min, max int) int {
	return rand.Intn(max-min) + min
}

// 生产队列
func producer(ch chan any, id int) {
	fmt.Printf("#%d 开始生产数据...\n", id)
	for i := 0; i < 5; i++ {
		ch <- rand.Intn(100)
	}
	// 500ms ~ 2s
	time.Sleep(time.Duration(randBetween(500, 2000)) * time.Millisecond)
	close(ch)
}

func fanOut(ch <-chan any, async bool, outs ...chan any) {
	go func() {
		defer func() { //退出时关闭所有的输出chan
			for i := 0; i < len(outs); i++ {
				close(outs[i])
			}
		}()
		for v := range ch { // 从输入chan中读取数据
			v := v
			for i := 0; i < len(outs); i++ {
				i := i
				if async {
					go func() {
						outs[i] <- v
					}()
				} else {
					outs[i] <- v // 放入到输出chan中，同步方式
				}
			}
		}
	}()
}

func main() {
	in := make(chan any)
	go producer(in, 1)

	out1 := make(chan any)
	out2 := make(chan any)

	fanOut(in, true, out1, out2)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range out1 {
			fmt.Println("Received from out1:", v)
		}
	}()

	go func() {
		defer wg.Done()
		for v1 := range out2 {
			fmt.Println("Received from out2:", v1)
		}
	}()
	wg.Wait()
}
