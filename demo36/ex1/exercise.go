package ex1

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// 练习1：完成以下代码，创建一个生产者，生产 1-10 的数字，
// 创建两个消费者，同时消费这些数字并打印出来。
// 要求：使用 WaitGroup 确保所有数据都被消费完
func exercise1() {
	numsChan := make(chan int)
	var wg sync.WaitGroup

	go func() {
		defer close(numsChan)
		for i := 1; i <= 10; i++ {
			numsChan <- i
		}
	}()

	// 两个消费者
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for num := range numsChan {
				fmt.Printf("消费者%d 正在处理数字：%d\n", id, num)
			}
		}(i)
	}
	wg.Wait()
}

// 练习2：实现一个工作池，要求：
// 1. 最多同时运行3个工作协程
// 2. 每个工作耗时1秒
// 3. 如果工作执行超过2秒，则超时退出
// 4. 使用channel来控制并发数量

func exercise2() {
	jobs := make(chan int, 10)
	results := make(chan string, 10)

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				select {
				case <-time.After(2 * time.Second):
					fmt.Printf("工作%d 超时退出\n", job)
				case results <- fmt.Sprintf("%s-%d:%d", "工作", workerID, job):
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	// 发送工作
	go func() {
		defer close(jobs)
		for i := 1; i <= 5; i++ {
			jobs <- i
		}
	}()

	// 如果用 for range 来接收结果，需要关闭 results channel ，否则会死锁
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("打印结果：%s\n", result)
	}
}

// 练习3：实现一个数据处理管道，要求：
// 1. 第一阶段：生成 1-100 的数字
// 2. 第二阶段：过滤出偶数
// 3. 第三阶段：将偶数乘以2
// 4. 可以通过 context 随时取消整个处理流程
// 5. 所有 goroutine 都应该正常退出，不能发生泄露
func exercise3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 第一阶段：生成 1-100 的数字
	numberChan := make(chan int)
	go func() {
		defer close(numberChan)
		for i := 1; i <= 100; i++ {
			select {
			case <-ctx.Done():
				return
			case numberChan <- i:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 第二阶段：过滤出偶数
	evenChan := make(chan int)
	go func() {
		defer close(evenChan)
		for num := range numberChan {
			if num%2 == 1 {
				continue
			}
			select {
			case <-ctx.Done():
				return
			case evenChan <- num:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 第三阶段：将偶数乘以2
	resultChan := make(chan int)
	go func() {
		defer close(resultChan)
		for num := range evenChan {
			select {
			case <-ctx.Done():
				return
			case resultChan <- num * 2:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 打印结果
	count := 0
	for result := range resultChan {
		fmt.Printf("最终结果：%d\n", result)
		count++
		if count > 10 {
			cancel()
			break
		}
	}
}

// 练习4：两个协程交替打印字符和数字

type Token struct{}

func exercise4() {
	numChan := make(chan Token)
	letterChan := make(chan Token)

	go func() {
		for {
			token := <-numChan
			fmt.Println(rand.Intn(100))
			time.Sleep(1 * time.Second)
			letterChan <- token
		}
	}()

	go func() {
		for {
			token := <-letterChan
			randIndex := rand.Intn(26)
			fmt.Printf("%c\n", 'A'+randIndex)
			time.Sleep(1 * time.Second)
			numChan <- token
		}
	}()
	numChan <- Token{}
	select {}
}
