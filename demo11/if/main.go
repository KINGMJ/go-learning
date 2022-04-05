package main

import (
	"fmt"
)

func main() {
	ifStatement()
}

/*
	if 作用域
*/
func ifStatement() {
	x := 0
	// n 所在的作用域为 if 语句，在外部不能访问
	if n := "abc"; x == 0 {
		fmt.Printf("%c\n", n[2]) // c
	} else if x < 0 {
		fmt.Println(n[1])
	} else {
		fmt.Printf("%c\n", n[0])
	}
	// 编译失败：undeclared name: n
	//fmt.Println(n)
}
