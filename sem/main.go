package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/semaphore"
)

func main() {
	var (
		maxWorker = 1
		sem       = semaphore.NewWeighted(int64(maxWorker))
		weight    = 1
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
	// time.Sleep(1 * time.Second)
}
