package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	demo2()
}

func demo1() {
	g := new(errgroup.Group)
	// 设置同时运行的 goroutine 的最大数量为 2。
	g.SetLimit(2)

	for i := 0; i < 5; i++ {
		i := i
		g.Go(func() error {
			time.Sleep(time.Second)
			fmt.Printf("任务：#%d 正在执行...\n", i)
			return nil
		})
	}

	if err := g.Wait(); err == nil {
		fmt.Println("Successful exec all")
	} else {
		fmt.Println("failed:", err)
	}
}

func demo2() {
	g := new(errgroup.Group)
	// 设置同时运行的 goroutine 的最大数量为 2。
	g.SetLimit(2)
	for i := 0; i < 5; i++ {
		i := i
		g.Go(func() error {
			time.Sleep(time.Second)
			fmt.Printf("任务：#%d 正在执行...\n", i)
			return nil
		})
	}
	ok := g.TryGo(func() error {
		time.Sleep(5 * time.Second)
		return nil
	})
	if !ok {
		fmt.Println("无空闲协程")
		return
	}

	if err := g.Wait(); err == nil {
		fmt.Println("Successful exec all")
	} else {
		fmt.Println("failed:", err)
	}
}
