package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func main() {
	demo2()
}

func demo1() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func write() {
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10ms
	rwlock.Unlock()
	defer wg.Done()
}

func read() {
	rwlock.RLock() // 加读锁
	time.Sleep(10 * time.Millisecond)
	rwlock.RUnlock()
	defer wg.Done()
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 一个线程安全的计数器
type Counter struct {
	mu    sync.RWMutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func demo2() {
	var counter Counter
	// 10个线程去读取
	for i := 0; i < 10; i++ {
		go func() {
			for {
				count := counter.Count() // 计数器读操作
				fmt.Println(count)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 每隔5s写一次
	for {
		counter.Incr() // 计数器写操作
		time.Sleep(5 * time.Second)
	}
}
