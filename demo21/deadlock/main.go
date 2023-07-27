package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 派出所证明
	var psCertificate sync.Mutex
	// 物业证明
	var propertyCertificate sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done() // 派出所处理完成
		psCertificate.Lock()
		defer psCertificate.Unlock()

		// 检查材料
		time.Sleep(5 * time.Second)
		// 请求物业证明
		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()
	}()

	go func() {
		defer wg.Done() // 物业处理完成
		propertyCertificate.Lock()
		defer propertyCertificate.Unlock()

		// 检查材料
		time.Sleep(5 * time.Second)
		// 请求派出所证明
		psCertificate.Lock()
		defer psCertificate.Unlock()
	}()

	wg.Wait()
	fmt.Println("成功完成")
}

/*
	死锁，互相等待
*/
