package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int64

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

/*
	两个线程并发去累加变量 x，导致会存在数据竞争。最后的结果与期望的不符
*/
