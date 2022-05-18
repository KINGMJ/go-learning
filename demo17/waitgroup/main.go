package main

import (
	"fmt"
	"sync"
)

func main() {
	waitGroupDemo()
}

func waitGroupDemo() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("goroutine 1 done!")
		defer wg.Done()
	}()

	go func() {
		fmt.Println("goroutine 2 done!")
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("all work done")
}
