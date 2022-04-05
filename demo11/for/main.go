package main

import (
	"fmt"
)

func main() {
	forStatement()
}

func forStatement() {
	// 在初始化语句中计算出全部结构是个好主意，len(s)
	// i, n 的作用域为 for 表达式
	s := "abc"

	// 常见的 for 循环，支持初始化语句。
	for i, n := 0, len(s); i < n; i++ {
		fmt.Printf("%c\n", s[i])
	}

	// 替代 while (n > 0) {}
	n := len(s) - 1
	for n > 0 {
		fmt.Printf("%c\n", s[n])
		n--
	}

	// 替代 while (true) {}
	for {
		fmt.Println(s)
	}
}
