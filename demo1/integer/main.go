package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {
	demo3()
}

/*
 整数类型
*/
func demo1() {
	// int 类型
	var i1 int = 12
	fmt.Println(i1)

	// 打印数据类型
	fmt.Println(reflect.TypeOf(i1)) // int

	// zero value 是 0
	var i2 int
	fmt.Println(i2) // 0
}

/*
 整数转字符串
*/
func intToString() {
	// 类型转换
	// Convert int to string with strconv.Itoa
	var i3 int = -38
	var s3 = strconv.Itoa(i3)
	fmt.Println(s3)                 // -38
	fmt.Println(reflect.TypeOf(s3)) // string

	// Convert int to string with fmt.Sprintf
	i4 := fmt.Sprintf("%d", i3)
	fmt.Println(i4)
}

/*
	int 类型之间的区别
*/
func demo3() {
	var i1 int = 1
	var i2 int8 = 127
	var i3 int16 = 32767
	fmt.Printf("size: %d, type: %s\n", unsafe.Sizeof(i1), reflect.TypeOf(i1))
	fmt.Printf("size: %d, type: %s\n", unsafe.Sizeof(i2), reflect.TypeOf(i2))
	fmt.Printf("size: %d, type: %s\n", unsafe.Sizeof(i3), reflect.TypeOf(i3))
}

/*
	1. 整数转字符串：strconv.Itoa 或 fmt.Sprintf
*/
