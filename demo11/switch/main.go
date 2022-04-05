package main

import (
	"fmt"
	"math/rand"
)

func main() {

}

func switchStatement() {
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
func switchStatement2() {
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

/*
	使用 switch 做类型断言
*/
func switchStatement3() {
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
