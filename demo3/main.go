package main

import (
	"fmt"
	"reflect"
	// "io"
	"strconv"
)

func main() {
	blankIdentifierDemo()
}

/*
  变量声明
*/
func declarVariableDemo() {
	var topLevel int64 = 5
	var (
		intVal int
		str    string = "str"
		fn     func(a int) string
	)

	intVal = 6

	fn = func(a int) string {
		str = "Hello World"
		fmt.Println(str)
		fmt.Println(intVal)
		return strconv.Itoa(a)
	}
	fn(int(topLevel))
}

/*
  多变量赋值
*/
func declarVariableDemo2() {
	// 类型推断
	var a = 123
	fmt.Println(reflect.TypeOf(a))

	// 多变量赋值
	var width, height int
	width = 40
	height = 50
	fmt.Println(width + height)

	// or
	var w, h int = 40, 50
	fmt.Println(w + h)

	// 不同类型，可以使用类型推断
	var w1, h1 = "40", 50
	fmt.Println(w1 + strconv.Itoa(h1)) // 4050
}

/*
  根据函数的返回值为变量赋值
*/
func declarVariableDemo3() {
	var multiReturn = func() (int, int) {
		return 1, 2
	}

	var multiReturn2 = func() (a int, b int) {
		a, b = 3, 4
		return
	}

	x, y := multiReturn()
	a, b := multiReturn2()

	fmt.Println(x + y) //3
	fmt.Println(a * b) //12
}

/*
	空白标识符
*/
func blankIdentifierDemo() {
	pets := []string{"dog", "cat", "fish"}
	// range 返回的是 index,value ，这里我们只需要 value
	for _, pet := range pets {
		fmt.Println(pet)
	}
}
