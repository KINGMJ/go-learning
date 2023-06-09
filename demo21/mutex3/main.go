package main

import (
	"fmt"
	"sync"
)

func main() {
	demo2()
}

type Counter struct {
	Name string
	sync.Mutex
	Count uint64
}

func demo1() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

type Counter2 struct {
	CounterType int
	Name        string
	sync.Mutex
	count uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *Counter2) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *Counter2) Count() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.count
}

func demo2() {
	var counter Counter2
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}
