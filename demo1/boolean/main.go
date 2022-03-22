package main

import (
	"fmt"
	"unsafe"
)

/*
	布尔类型
*/
func main() {
	// bool 类型，true or false
	var b bool = true
	fmt.Println(b)
	// zero value 是 false
	var b2 bool
	fmt.Println(b2) // false

	// bool 的 size 是 1
	fmt.Println(unsafe.Sizeof(b2)) // 1
}
