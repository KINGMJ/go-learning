package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))              // 3
	fmt.Println(add("hello ", "world")) // hello world
}

type Addable interface {
	int | string
}

func add[T Addable](a, b T) T {
	return a + b
}
