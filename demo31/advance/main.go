package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有取消功能的 context 和对应的 cancel 函数
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "name", "jack")

	go longRunningOperation(ctx)
	time.Sleep(2 * time.Second)
	// 模拟运行一段时间后取消操作
	cancel()
	// 等待一段时间，以观察操作是否已经取消
	time.Sleep(1 * time.Second)
}

func longRunningOperation(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		value, _ := ctx.Value("name").(string)
		fmt.Println(value)

	}
}
