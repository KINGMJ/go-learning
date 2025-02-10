package ex2

import (
	"sync"
	"time"
)

// 练习2：实现一个并发安全的缓存，要求：
// 1. 支持读写操作
// 2. 使用读写锁提高并发性
// 3. 支持过期时间
// 4. 支持最大容量限制
type Cache interface {
	Set(key string, value any, exp time.Duration)
	Get(key string) (any, bool)
	Delete(key string)
	Clear()
	Len() int
	Data() map[string]item
}

type cache struct {
	data map[string]item
	mu   sync.RWMutex
	// 添加自动清理时间间隔
	cleanupInterval time.Duration
}

type item struct {
	value any
	exp   time.Time
}

func NewCache() Cache {
	c := &cache{
		data:            make(map[string]item),
		cleanupInterval: 10 * time.Second,
	}
	go c.cleanupLoop()
	return c
}

func (c *cache) cleanupLoop() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()
	for range ticker.C {
		c.cleanup()
	}
}

func (c *cache) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for key, item := range c.data {
		if now.After(item.exp) {
			delete(c.data, key)
		}
	}
}

func (c *cache) Set(key string, value any, exp time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expiration := time.Now().Add(exp)
	c.data[key] = item{
		value: value,
		exp:   expiration,
	}
}

func (c *cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.data[key]
	if !ok {
		return nil, false
	}
	if time.Now().After(item.exp) {
		return nil, false
	}
	return item.value, true
}

func (c *cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]item)
}

func (c *cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

func (c *cache) Data() map[string]item {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data
}
