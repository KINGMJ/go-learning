package main

import (
	"fmt"
)

func main() {
	declarIota()
}

/*
	声明 iota 常量
*/
func declarIota() {
	const (
		Low = iota // 0
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
func declarIota2() {
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
func declarIota3() {
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
