package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	goroutineDemo2()
}

func goroutineDemo() {
	go hello()
	fmt.Println("main goroutine done")
	time.Sleep(time.Second)

	// main goroutine done
	// Hello Goroutinex!
}
func hello() {
	fmt.Println("Hello Goroutine!")
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

var wg sync.WaitGroup

func routine(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func goroutineDemo2() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go routine(i)
	}
	wg.Wait()
}
