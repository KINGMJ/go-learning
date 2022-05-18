package main

import (
	"fmt"
	"sync"
	"time"
)

var cache map[int]int
var mu sync.Mutex

func expensiveOperation(n int) int {
	// in real code this operation would be very expensive
	time.Sleep(1 * time.Second)
	return n
}

func getCached(n int) int {
	mu.Lock()
	v, isCached := cache[n]
	mu.Unlock()
	if isCached {
		return v
	}

	v = expensiveOperation(n)

	mu.Lock()
	cache[n] = v
	mu.Unlock()
	return v
}

func accessCache() {
	total := 0
	for i := 0; i < 5; i++ {
		n := getCached(i)
		total += n
	}
	fmt.Printf("total: %d\n", total)
}

func main() {
	cache = make(map[int]int)
	go accessCache()
	accessCache()
}
