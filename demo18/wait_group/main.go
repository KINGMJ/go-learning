package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	demo3()
}

func demo1() {
	wg.Add(10)
	wg.Add(-10)
	wg.Add(-1)
}

func demo2() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()
	wg.Done()
}

func demo3() {
	var wg sync.WaitGroup
	go domeSomething(1000, &wg)
	go domeSomething(2000, &wg)
	wg.Wait()
	fmt.Print("Done")
}

func domeSomething(millisecs time.Duration, wg *sync.WaitGroup) {
	duration := millisecs * time.Millisecond
	time.Sleep(duration)
	wg.Add(1)
	fmt.Println("后台执行，duration：", duration)
	wg.Done()
}
