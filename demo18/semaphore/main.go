package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	demo1()
}

func demo1() {
	var (
		maxWorker = 10
		sem       = semaphore.NewWeighted(int64(maxWorker))
		weight    = 2
	)

	urls := []string{
		"http://www.example.com",
		"http://www.example.net",
		"http://www.example.net/foo",
		"http://www.example.net/bar",
		"http://www.example.net/baz",
		"http://www.baidu.com",
		"http://www.test.com",
	}

	for _, u := range urls {
		sem.Acquire(context.Background(), int64(weight))
		go func(u string) {
			doSomething(u)
			sem.Release(int64(weight))
		}(u)
	}
	fmt.Println("All Done...")
}

func doSomething(url string) {
	fmt.Printf("抓取url：%s\n", url)
	time.Sleep(1 * time.Second)
}

func demo2() {
	var (
		maxWorkers       = runtime.GOMAXPROCS(0)                    // worker 数量
		sema             = semaphore.NewWeighted(int64(maxWorkers)) // 信号量
		task             = make([]int, maxWorkers*4)                // 任务数，是 worker 的四倍
		weight     int64 = 2
	)

	ctx := context.Background()

	for i := range task {
		// 如果没有worker可用，会阻塞在这里，直到某个worker被释放
		if err := sema.Acquire(ctx, weight); err != nil {
			break
		}
		// 启动 worker goroutine
		go func(i int) {
			defer sema.Release(weight)
			time.Sleep(1 * time.Second) // 模拟一个耗时操作
			fmt.Printf("处理数据：%d\n", i)
			task[i] = i + 1
		}(i)
	}
	// 请求所有的worker,这样能确保前面的worker都执行完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("获取所有的worker失败：%v", err)
	}
	fmt.Println(task)
}
