package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	demo2()
}

func demo1() {
	var group singleflight.Group
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		// 启动 5个goroutine
		go func(id int) {
			defer wg.Done()
			// Do函数是我们要执行的操作，可以是一些需要耗时的操作，比如请求外部服务等
			result, err, _ := group.Do("key", func() (interface{}, error) {
				fmt.Printf("Executing operation for ID: #%d\n", id)
				time.Sleep(1 * time.Second)
				return fmt.Sprintf("Result is %d", id), nil
			})
			if err != nil {
				fmt.Printf("Error for ID %d: %v\n", id, err)
				return
			}
			fmt.Printf("Result for ID: #%d: %v\n", id, result)
		}(i)
	}
	wg.Wait()
}

func demo2() {
	var group singleflight.Group
	var wg sync.WaitGroup
	resultChans := make([]<-chan singleflight.Result, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 使用 DoChan 异步获取结果
			resultChans[id] = group.DoChan("key", func() (interface{}, error) {
				fmt.Printf("Executing operation for ID: #%d\n", id)
				time.Sleep(1 * time.Second)
				return fmt.Sprintf("Result is %d", id), nil
			})
		}(i)
	}

	wg.Wait()

	fmt.Println("Canceling the waiting result...")
	group.Forget("key")
	// 在这里等待所有结果
	for id, resultChan := range resultChans {
		result := <-resultChan
		// 根据 Result 中的 Shared 字段判断结果是否是被共享的
		if result.Shared {
			fmt.Printf("Got shared result for ID #%d: %v\n", id, result.Val)
		} else {
			fmt.Printf("Got original result for ID #%d: %v\n", id, result.Val)
		}
	}
}
