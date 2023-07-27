package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

func main() {
	demo3()
}

/*
go 不是可重入锁，执行失败
*/
func demo1() {
	var foo func(sync.Locker)
	var bar func(sync.Locker)

	foo = func(l sync.Locker) {
		fmt.Println("in foo")
		l.Lock()
		bar(l)
		l.Unlock()
	}

	bar = func(l sync.Locker) {
		l.Lock()
		fmt.Println("in bar")
		l.Unlock()
	}
	l := &sync.Mutex{}
	foo(l)
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

/*
实现可重入锁
*/

var mu RecursiveMutex

func demo2() {
	var foo func()
	var bar func()

	foo = func() {
		fmt.Println("in foo")
		mu.Lock()
		bar()
		mu.Unlock()
	}

	bar = func() {
		mu.Lock()
		fmt.Println("in bar")
		mu.Unlock()
	}
	foo()
}

type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的 goroutine id
	recursion int32 // 这个goroutine 重入的次数
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	// 如果这个goroutine还没有完全释放，则直接返回
	if m.recursion != 0 {
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

/*
实现可重入锁
*/

var tokenMu TokenRecursiveMutex
var wg sync.WaitGroup
var x int64

func demo3() {
	wg.Add(2)
	go add(1000)
	go add(2000)
	wg.Wait()
	fmt.Println(x) // 10000
}

func add(token int64) {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		tokenMu.Lock(token)
		x = x + 1
		tokenMu.Unlock(token)
	}
}

type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int32
}

func (m *TokenRecursiveMutex) Lock(token int64) {
	//如果传入的token和持有锁的token一致，说明是递归调用
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

func (m *TokenRecursiveMutex) Unlock(token int64) {
	// 释放其它token持有的锁
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.token, token))
	}
	m.recursion--
	// 还没有回退到最初的递归调用
	if m.recursion != 0 {
		return
	}
	atomic.StoreInt64(&m.token, 0) // 没有递归调用了，释放锁
	m.Mutex.Unlock()
}
