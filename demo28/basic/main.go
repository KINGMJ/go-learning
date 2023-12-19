package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func main() {
	demo7()
}

var x int64
var mutex sync.Mutex
var wg sync.WaitGroup

func add() {
	x++
	wg.Done()
}

func mutexAdd() {
	mutex.Lock()
	x++
	mutex.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

// 各种操作的效率对比
func demo1() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go atomicAdd()
	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Since(start))
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func demo2() {
	var n atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			n.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(n.Load())
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 实现原子减法操作
func demo3() {
	var (
		m uint64 = 2
		n uint64 = 129
		k int    = 2
	)
	show := fmt.Println
	atomic.AddUint64(&n, -m)
	show(n)
	atomic.AddUint64(&n, ^(m - 1))
	show(n)
	atomic.AddUint64(&n, -uint64(k))
	show(n)
	atomic.AddUint64(&n, ^uint64(k-1))
	show(n)
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 交换操作
func demo4() {
	var n atomic.Int64
	n.Store(123)
	// 交换操作
	var old = n.Swap(789)
	fmt.Println(n.Load(), old) // 789 123

	swapped := n.CompareAndSwap(123, 456)
	fmt.Println(swapped)  // false
	fmt.Println(n.Load()) // 789

	swapped = n.CompareAndSwap(789, 456)
	fmt.Println(swapped)  // true
	fmt.Println(n.Load()) // 456
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 指针类型的原子操作

type T struct{ x int }

func demo5() {
	var pt *T
	// 这里转换成 *unsafe.Pointer 类型，因为 StorePointer 第一个参数为该类型
	var unsafePPT = (*unsafe.Pointer)(unsafe.Pointer(&pt))
	var ta, tb = T{1}, T{2}
	// 修改
	atomic.StorePointer(unsafePPT, unsafe.Pointer(&ta))
	fmt.Println(pt) // &{1}

	// 读取
	pa1 := (*T)(atomic.LoadPointer(unsafePPT))
	fmt.Println(pa1 == &ta) // true

	// 置换
	pa2 := (atomic.SwapPointer(unsafePPT, unsafe.Pointer(&tb)))
	fmt.Println((*T)(pa2) == &ta) // true
	fmt.Println(pt)               // &{2}

	// 比较置换
	b := atomic.CompareAndSwapPointer(unsafePPT, pa2, unsafe.Pointer(&tb))
	fmt.Println(b) // false

	b = atomic.CompareAndSwapPointer(unsafePPT, unsafe.Pointer(&tb), pa2)
	fmt.Println(b) // true
}

// 使用泛型实现
func demo6() {
	var pt atomic.Pointer[T]
	var ta, tb = T{1}, T{2}
	// 修改
	pt.Store(&ta)
	fmt.Println(pt.Load()) // &{1}
	// 读取
	pa1 := pt.Load()
	fmt.Println(pa1 == &ta) // true
	// 交换
	pa2 := pt.Swap(&tb)
	fmt.Println(pa2 == &ta) // true
	fmt.Println(pt.Load())  // &{2}

	b := pt.CompareAndSwap(&ta, &tb)
	fmt.Println(b) // false
	b = pt.CompareAndSwap(&tb, &ta)
	fmt.Println(b) // true
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 任何类型值的原子操作
func demo7() {
	type T struct{ a, b, c interface{} }
	var ta = T{1, "hello", 1.2}
	var v atomic.Value
	v.Store(ta)
	var tb = v.Load().(T)
	fmt.Println(tb) // {1 hello 1.2}

	fmt.Println(ta == tb) // true

	v.Store("hello") // 导致 panic
}
