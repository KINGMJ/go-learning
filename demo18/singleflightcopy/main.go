package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	var group singleflight.Group
	var mu sync.Mutex

	// 查询缓存的函数
	queryCache := func(key string) (interface{}, error) {
		// 模拟查询缓存
		fmt.Printf("Querying cache for key: %s\n", key)
		time.Sleep(1 * time.Second)
		return fmt.Sprintf("Data for key %s", key), nil
	}

	// 获取数据的函数，使用 singleflight 防止缓存击穿
	getData := func(key string) (interface{}, error) {
		fmt.Printf("Fetching data for key: %s\n", key)
		// 使用 singleflight.Do 来执行查询
		result, err, _ := group.Do(key, func() (interface{}, error) {
			// 查询缓存
			data, err := queryCache(key)
			if err != nil {
				return nil, err
			}
			// 假设这里有一些处理数据的逻辑
			return data, nil
		})
		return result, err
	}

	// 并发执行获取数据的操作
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		key := fmt.Sprintf("key%d", i)
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			// 获取数据
			data, err := getData(k)
			mu.Lock()
			defer mu.Unlock()
			if err == nil {
				fmt.Printf("Received data: %v\n", data)
			} else {
				fmt.Printf("Error fetching data for key %s: %v\n", k, err)
			}
		}(key)
	}

	// 假设在一段时间后我们不再关心某个查询
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Canceling the waiting query for key2...")
	// 	group.Forget("key2")
	// }()

	wg.Wait()
}
