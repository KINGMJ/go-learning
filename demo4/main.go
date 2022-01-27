package main

import (
	"fmt"
	"reflect"
)

func main() {
	iotaDemo3()
}

func constantDemo() {
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

func constantDemo2() {
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

func iotaDemo() {
	const (
		Low = iota
		Medium
		High
	)
	//Low: 0
	//Medium: 1
	//High: 2
	fmt.Printf("Low: %d\nMedium: %d\nHigh: %d\n", Low, Medium, High)

	//iota 可以用来创建位掩码，相当于 1 << 0，1 << 1，1 << 2
	const (
		Secure = 1 << iota // 0b001
		Authn              // 0b010
		Ready              // 0b100
	)
	ConnState := Secure | Authn
	fmt.Printf(`Secure:    0x%x (0b%03b)
Authn:     0x%x (0b%03b)
ConnState: 0x%x (0b%03b)
`, Secure, Secure, Authn, Authn, ConnState, ConnState)
}

/*
	iota 跳过值
*/
func iotaDemo2() {
	const ( // iota 重置为 0
		a = 1 << iota // a == 1
		b = 1 << iota // b == 2
		c = 3         // c == 3（iota 未使用但仍递增）
		d = 1 << iota // d == 8
	)
	fmt.Printf("a: %d, b: %d, c: %d, d: %d\n", a, b, c, d)

	const (
		a1 = iota // a1 = 0
		_         // iota 递增到 1
		b1        // b1 = 2
	)
	fmt.Printf("a1: %d, b1: %d\n", a1, b1)
}

/*
	iota 在表达式中使用
*/
func iotaDemo3() {
	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
		bit1, mask1                          // bit1 == 2, mask1 == 1
		_, _                                 // 跳过 iota == 2
		bit3, mask3                          // bit3 == 8, mask3 == 7
	)
	fmt.Printf("bit0: %d, mask0: 0x%x\n", bit0, mask0)
	fmt.Printf("bit3: %d, mask3: 0x%x\n", bit3, mask3)

	// 表示字节
	type ByteSize int
	const (
		_           = iota // 通过分配给空白标识符来忽略第一个值
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)
	fmt.Printf("KB: %d\n", KB)
}
