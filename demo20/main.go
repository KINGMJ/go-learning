package main

import "fmt"

func main() {
	panicInDefer()
}

// panic 例子
func panicDemo() {
	defer fmt.Println("Exiting main")
	panicBar()
	fmt.Println("不会执行")
}
func panicBar() {
	defer fmt.Println("Exiting foo")
	panic("bar")
	fmt.Println("不会执行")
	defer fmt.Println("会执行吗")
}

// recover 例子
func foo() {
	panic("bar")
}

func bar() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Printf("Recovered with message %s\n", msg)
		}
	}()
	foo()
	fmt.Println("Never gets executed")
}

func recoverDemo() {
	fmt.Println("Entering main")
	bar()
	fmt.Println("Exiting main the normal way")
}

// recover 例子2
func recoverDemo2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err.(string))
		}
	}()
	panic("panic error!")
}

// panic 与 recover 的使用
func panicAndRecoverDemo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
	// send on closed channel
}

// defer 中使用 panic
func panicInDefer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("defer panic1")
	}()

	defer func() {
		panic("defer panic2")
	}()

	panic("normal panic")
	// defer panic1
}
