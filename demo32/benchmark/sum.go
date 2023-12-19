package benchmark

import (
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mu sync.Mutex

func MutexSum() int64 {
	var count int64 = 0
	for i := 0; i < 100000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
	return count
}

func AtomicSum() int64 {
	var count int64 = 0
	for i := 0; i < 100000; i++ {
		atomic.AddInt64(&count, 1)
	}
	return count
}
