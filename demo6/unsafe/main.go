package main

import (
	"fmt"
	"unsafe"
)

func main() {
	demo2()
}

func demo1() {
	var t1 float64 = 31211
	fmt.Println(Float64bits(t1))
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

type User struct {
	Name string
	Age  int
}

func demo2() {
	u := new(User)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "张三"

	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.Age)))
	*pAge = 20
	fmt.Println(*u)
}
