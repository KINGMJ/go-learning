package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int // job id
	RandNum int // 需要计算的随机数
}

type Result struct {
	job *Job // job 实例
	sum int  // 求和
}

func main() {
	// 任务数为 5
	const numJobs = 5
	jobs := make(chan *Job, numJobs)      // job 管道
	result := make(chan *Result, numJobs) // result 管道

	// 创建工作池来并发处理 100 个任务
	createPool(3, jobs, result)

	// 开个打印的协程
	go func(result chan *Result) {
		// 遍历结果管道打印
		for result := range result {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(result)

	// 循环创建job，输入到管道
	var id int
	for j := 1; j <= 100; j++ {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobs <- job
	}
}

/*
	创建工作池
	num 工作线程数
*/
func createPool(num int, jobs <-chan *Job, result chan<- *Result) {
	for i := 1; i <= num; i++ {
		// 开启一个线程
		go func(jobs <-chan *Job, result chan<- *Result) {
			// 执行任务
			// 遍历 job 管道所有数据，进行相加
			for job := range jobs {
				num := job.RandNum
				var sum int
				for num != 0 {
					tmp := num % 10
					sum += tmp
					num /= 10
				}

				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				result <- r
			}
		}(jobs, result)
	}
}
