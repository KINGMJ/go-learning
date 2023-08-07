package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	queue := NewSliceQueue(10)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			queue.Enqueue(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(queue.data...)
}

type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, 0, n)}
}

// 进队列
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}
