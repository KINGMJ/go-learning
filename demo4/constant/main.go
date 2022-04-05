package main

import (
	"fmt"
	"reflect"
)

func main() {
	declarConstant()
}

/*
	常量声明
*/
func declarConstant() {
	// 单行常量声明
	const greeting string = "Hello World"

	// 多行常量声明
	const (
		int1    int = 32
		string1     = "string"
		int2        = 33
	)

	fmt.Println(greeting)
	fmt.Println(int1)
}

/*
	无类型常量
*/
func declarConstant2() {
	// untypedNumber 因为是常量，所以它保持无类型，直到它被分配给一个变量
	const untypedNumber = 345
	fmt.Println(reflect.TypeOf(untypedNumber)) // int

	// 无需转换为 uint16
	var u16 uint16 = untypedNumber
	fmt.Println(reflect.TypeOf(u16)) // uint16

	// 无需转换为 float64
	var f float64 = untypedNumber
	fmt.Println(reflect.TypeOf(f))

	// var b int8 = untypedNumber // 报错，编译器检测到345太大无法放入int8
}
