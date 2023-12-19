package main

import (
	"fmt"
	"sync"
)

type MyObj struct {
	Value int
}

func main() {
	objectPool := &sync.Pool{
		New: func() any {
			fmt.Println("Creating a new Object.")
			return &MyObj{}
		},
	}

	obj1 := objectPool.Get().(*MyObj)
	obj1.Value = 1

	objectPool.Put(2)
	obj2 := objectPool.Get()
	fmt.Println(obj1.Value)
	fmt.Println(obj2)
}
