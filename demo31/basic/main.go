package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	demo1()
}

func demo1() {
	// 创建一个带有取消功能的 context 和对应的 cancel 函数
	ctx, cancel := context.WithCancel(context.Background())
	// 启动一个长时间运行的操作
	go longRunningOperation(ctx)
	time.Sleep(2 * time.Second)
	// 模拟运行一段时间后取消操作
	cancel()
	// 等待一段时间，以观察操作是否已经取消
	time.Sleep(1 * time.Second)
}

func longRunningOperation(ctx context.Context) {
	for {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Working...")
		case <-ctx.Done():
			fmt.Println("Operation canceled.")
			return
		}
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
func demo2() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	operationWithDeadline(ctx)
}

func operationWithDeadline(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("Operation deadline:", deadline)
	}
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Operation completed.")
	case <-ctx.Done():
		fmt.Println("Operation canceled due to deadline.")
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
func demo3() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	operationWithTimeout(ctx)
}

func operationWithTimeout(ctx context.Context) {
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Operation completed.")
	case <-ctx.Done():
		fmt.Println("Operation canceled due to timeout.")
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
func demo4() {
	// 设置一个截止日期为相对的2秒钟
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 将截止日期存储在 context 中
	ctx = context.WithValue(ctx, "timeout", 1*time.Second)
	operationWithValue(ctx)
}

func operationWithValue(ctx context.Context) {
	// 从 context 中获取传递值
	timeout, ok := ctx.Value("timeout").(time.Duration)
	fmt.Println(timeout)
	if !ok {
		fmt.Println("Timeout not set.")
		return
	}
	select {
	case <-time.After(timeout):
		fmt.Println("Operation completed.")
	case <-ctx.Done():
		fmt.Println("Operation canceled due to timeout.")
	}
}
