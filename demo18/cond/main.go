package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

var status int64

func main() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}
	time.Sleep(5 * time.Second)
	go broadcast(c)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	status = 1
	c.L.Unlock()
	// 通知条件变化
	c.Broadcast()
	// c.Signal()
}

func listen(c *sync.Cond) {
	c.L.Lock()
	// 等待满足条件
	for status != 1 {
		c.Wait()
	}
	fmt.Println("listen")
	c.L.Unlock()
}
