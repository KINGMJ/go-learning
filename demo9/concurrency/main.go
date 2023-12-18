package main

import (
	"fmt"
	"sync"
)

func main() {
	demo2()
}

func demo1() {
	var m = make(map[int]int)
	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
		}
	}()

	select {}
}

// 实现一个线程安全的 map

type RWMap[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V
}

func NewRWMap[K comparable, V any](n int) *RWMap[K, V] {
	return &RWMap[K, V]{
		m: make(map[K]V, n),
	}
}

func (m *RWMap[K, V]) Get(k K) (V, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

func (m *RWMap[K, V]) Set(k K, v V) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *RWMap[K, V]) Delete(k K) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RWMap[K, V]) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap[K, V]) Each(f func(k K, v V) bool) {
	m.RLock()
	defer m.RUnlock()
	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}

func demo2() {
	var m = NewRWMap[int, string](10)
	go func() {
		m.Set(1, "jack")
	}()

	go func() {
		for {
			val, _ := m.Get(1)
			fmt.Println(val)
		}
	}()

	select {}
}
