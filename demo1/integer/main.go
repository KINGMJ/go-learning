package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	intToString()
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
	1. 整数转字符串：strconv.Itoa 或 fmt.Sprintf
*/
