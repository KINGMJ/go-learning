package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	semaphoreSize      = 4
	mu                 sync.Mutex // 互斥锁
	totalTasks         int
	curConcurrentTasks int
	maxConcurrentTasks int
)

/*
	密集型计算
*/
func timeConsumingTask() {
	mu.Lock()
	totalTasks++
	curConcurrentTasks++
	// 限制最大的并发梳数量

	fmt.Println(curConcurrentTasks)

	if curConcurrentTasks > maxConcurrentTasks {
		maxConcurrentTasks = curConcurrentTasks
	}
	mu.Unlock()

	// 在实际系统中，这将是一个 CPU 密集型操作
	time.Sleep(10 * time.Millisecond)

	mu.Lock()
	curConcurrentTasks--
	mu.Unlock()
}

func main() {
	sem := make(chan struct{}, semaphoreSize)
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		sem <- struct{}{}
		wg.Add(1)
		go func() {
			timeConsumingTask()
			<-sem
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("total tasks         : %d\n", totalTasks)
	fmt.Printf("max concurrent tasks: %d\n", maxConcurrentTasks)
}

/*
	使用信号量来限制并发
*/
