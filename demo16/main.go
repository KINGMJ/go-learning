package main

import (
	"fmt"
)

func main() {
	deferDemo4()
}

func logExit(name string) {
	fmt.Printf("Function %s returned\n", name)
}

func deferDemo() {
	defer logExit("main")
	fmt.Println("First main statement")
	fmt.Println("Last main statement")
	// First main statement
	// Last main statement
	// Function main returned
}

func deferDemo2() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i) // 4 3 2 1 0
	}
}

func deferDemo3(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")

	defer func() {
		fmt.Println(100 / x)
	}()

	defer fmt.Println("c")
	//c
	//b
	//a
	//panic: runtime error: integer divide by zero
}

func deferDemo4() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func(i2 int) {
			fmt.Println(i2) // 4 3 2 1 0
		}(i)
	}
}
