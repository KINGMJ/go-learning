package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked      = 1 << iota // 加锁标志位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识 waiter 的起始 bit 位置
)

// 扩展一个 Mutex 结构
type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}
}
