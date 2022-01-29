package main

import "fmt"

func main() {
	arrayDemo()
}

func arrayDemo() {
	// 声明数组
	var arr [4]int
	// 数组的 zero-value 是每个元素的 zero-value
	fmt.Println(arr)         // [0 0 0 0]
	fmt.Printf("%#v\n", arr) // [4]int{0,0,0,0}

	arr = [4]int{1, 2, 3, 4}
	fmt.Printf("%#v\n", arr) // [4]int{1, 2, 3, 4}
}

func arrayDemo2() {
	arr := [2]int{4, 5}
	// 访问数组元素
	fmt.Println(arr[1]) // 5
	// 设置元素的值
	arr[1] = 3
	fmt.Println(arr) // [4 3]
	// 获取数组的长度
	fmt.Println(len(arr)) // 2

	// 使用 ... 替代数组的长度，可以根据初始值推断出数组的长度
	arr2 := [...]int{10, 20, 30, 40, 70}
	fmt.Println(arr2)
}

/*
	多维数组
*/
func multidimensionalArrayDemo() {
	multiDimArray := [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	// 可以简写成如下方式
	var simplified = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(multiDimArray)
	fmt.Println(simplified)
}
