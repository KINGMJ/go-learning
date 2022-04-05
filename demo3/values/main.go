package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("go" + "lang")

	// 结果相同，但类型不一样
	fmt.Println(reflect.TypeOf(1.1 + 1.9)) // float64
	fmt.Println(reflect.TypeOf(1 + 2))     // int
	fmt.Println(true && false)             // false

}

/*
	go 求值
*/
