package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	GoroutineDemo2()
}

func hello() {
	fmt.Println("Hello Goroutine!")
}

/*
	启动单个 goroutine
*/
func GoroutineDemo() {
	go hello()
	fmt.Println("main goroutine done")
	time.Sleep(time.Second)
	// main goroutine done
	// Hello Goroutinex!
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

var wg sync.WaitGroup

func routine(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func GoroutineDemo2() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go routine(i)
	}
	wg.Wait()
}
