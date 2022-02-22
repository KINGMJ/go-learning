package main

import (
	"fmt"
	"math/rand"
)

func main() {
	breakDemo()
}

func ifDemo() {
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

func switchDemo() {
	a := 1
	switch a {
	// case 支持多个值
	case 1, 3:
		fmt.Println("a is 1 or 3")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Printf("default is %d\n", a)
	}

	var n = 0
	// switch 省略表达式，相当于 if...else
	switch {
	case n > 0 && n < 10:
		fmt.Println("n > 0 and < 10")
	case n > 10 && n < 20:
		fmt.Println("n> 10 and <20")
	default:
		fmt.Println("def")
	}
}

/*
	switch 语句中声明变量
*/
func switchDemo2() {
	switch n := rand.Intn(9); n {
	case 1, 2, 3:
		fmt.Printf("case 1, 2, 3: n is %d\n", n)
	case 4, 5:
		fmt.Printf("case 4, 5: n is %d\n", n)
	default:
		fmt.Printf("default: n is %d\n", n)
	}

	// 编译失败，n 的作用域为 switch 语句
	//fmt.Println(n)
}

func switchDemo3() {
	var printType = func(iv interface{}) {
		switch v := iv.(type) {
		case int:
			fmt.Printf("'%d' is of type int\n", v)
		case string:
			fmt.Printf("'%s' is of type string\n", v)
		case float64:
			fmt.Printf("'%f' is of type float64\n", v)
		default:
			fmt.Printf("We don't support type '%T'\n", v)
		}
	}

	printType("5")
	printType(4)
	printType(true)
	//'5' is of type string
	//'4' is of type int
	//We don't support type 'bool'
}

func forDemo() {
	// 在初始化语句中计算出全部结构是个好主意，len(s)
	// i, n 的作用域为 for 表达式
	s := "abc"
	for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
		fmt.Printf("%c\n", s[i])
	}

	n := len(s) - 1
	for n > 0 { // 替代 while (n > 0) {}
		fmt.Printf("%c\n", s[n])
		n--
	}

	for { // 替代 while (true) {}
		fmt.Println(s)
	}
}

func gotoDemo() {
	var printIsOdd = func(n int) {
		if n%2 == 1 {
			goto isOdd
		}
		fmt.Printf("%d is even\n", n)
		// 需要结束掉，不然会执行 label 的语句
		return

		// label 语法
	isOdd:
		fmt.Printf("%d is odd\n", n)
	}

	printIsOdd(5)  // 5 is odd
	printIsOdd(10) // 10 is even
}

func breakDemo() {
	j := 100

loop:
	for j < 110 {
		j++

		fmt.Println(j)
		switch j % 3 {
		case 0:
			continue loop
		case 1:
			break loop
		}

		fmt.Println("Var : ", j)
	}
}
