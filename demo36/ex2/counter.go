package ex2

import "sync"

// 练习1：实现一个并发安全的计数器，要求：
// 1. 支持增加和获取操作
// 2. 保证并发安全
// 3. 提供批量操作方法
type Counter interface {
	Increament() int
	Get() int
	Add(delta int) int
	Reset()
}

type counter struct {
	value int
	mu    sync.Mutex
}

func NewCounter() Counter {
	return &counter{}
}

func (c *counter) Increament() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	return c.value
}

func (c *counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (c *counter) Add(delta int) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
	return c.value
}

func (c *counter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = 0
}
