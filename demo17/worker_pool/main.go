package main

import (
	"fmt"
	"time"
)

func main() {
	// 任务数为 5
	const numJobs = 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	// 开启 3 个线程
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-result
	}
}

// jobs 是一个接收通道，只能从通道接收值
// result 是一个发送通道，只能发送值到通道中
// 该函数的作用是从 jobs 里面接收数据，执行任务，将结果发送到 result 通道
func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		// 模拟工作执行
		time.Sleep(5 * time.Second)
		fmt.Println("worker", id, "finished job", j)
		// 将结果发送到 result
		result <- j * 2
	}
}
