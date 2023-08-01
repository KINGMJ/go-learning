package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	demo2()
}

func demo1() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		time.Sleep(time.Second * 5)
		fmt.Println("hello")
	}
}

func demo2() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	fmt.Println(runtime.GOOS)
	time.Sleep(time.Second)
}
