package main

import (
	"fmt"
)

func main() {
	twoDSlice()
}

/*
slice 声明
*/
func declareSlice() {
	// 声明一个 slice
	slice := make([]int, 0, 5)
	// 追加元素
	slice = append(slice, 5)
	// 追加多个元素
	slice = append(slice, 6, 7)
	fmt.Println(slice) // [5 6 7]
	// 切片的长度，就是当前元素的个数
	fmt.Println(len(slice)) // 3
	// 切片的容量，表示一个切片拥有的总元素数。这就是底层数组的大小
	fmt.Println(cap(slice)) //5
}

/*
nil slice
*/
func nilSlice() {
	s := []int(nil)
	if len(s) == 0 {
		fmt.Printf("s  is empty: %#v\n", s) // s  is empty: []int(nil)
	}

	var s2 []int
	if len(s2) == 0 {
		fmt.Printf("s2 is empty: %#v\n", s2) // s2 is empty: []int(nil)
	}

	// s3 是 empty slice 而不是 nil slice
	s3 := make([]int, 0)
	if len(s3) == 0 {
		fmt.Printf("s3 is empty: %#v\n", s3) // s3 is empty: []int{}
	}
}

/*
创建一个 slice
*/
func declareSlice2() {
	var a []int
	fmt.Printf("a is %#v\n", a) // a is []int(nil)
	a = append(a, 3)
	fmt.Printf("a is %#v\n", a) // a is []int(3)

	// 初始化一个 slice，是 nil slice
	var nilSlice []bool
	fmt.Printf("nilSlice is %#v\n", nilSlice) // nilSlice is []bool(nil)

	// 使用字面量创建一个空 slice
	empty1 := []bool{}
	fmt.Printf("empty1 is %#v\n", empty1) // empty1 is []bool{}
	empty1 = append(empty1, false)
	// slice 的 cap 是动态的，如果一个元素都没有时 0；插入了一个元素后是 8；超过 8 个元素，就变成 16
	fmt.Println(cap(empty1))

	// 使用 make 语法创建一个空 slice
	empty2 := make([]bool, 0)
	fmt.Printf("empty2 is %#v\n", empty2) // empty2 is []bool{}
	empty2 = append(empty2, true)
	fmt.Println(len(empty2)) // 1
	fmt.Println(cap(empty2)) // 8
}

/*
创建静态预分配切片
*/
func declareSlice3() {
	a := []int{1, 2, 3, 4}
	//b := make([]string, 4)

	fmt.Println(len(a)) // 4
	fmt.Println(cap(a)) // 4

	a = append(a, 1)
	fmt.Println(len(a)) // 5
	fmt.Println(cap(a)) // 8
}

/*
分配预期大小
*/
func declareSlice4() {
	// make 第二个参数是切片的长度，第三个参数是容量
	a := make([]int, 0, 12)
	fmt.Printf("a has length %d and capacity %d\n", len(a), cap(a)) // a has length 0 and capacity 12
	for i := 0; i < 16; i++ {
		a = append(a, i)
	}
	fmt.Printf("a has length %d and capacity %d\n", len(a), cap(a)) // a has length 16 and capacity 24
}

/*
slice 初始化
*/
func initSlice() {
	// 没有分配空间，直接赋值会报错
	var a []int
	// a[0] = 1
	fmt.Println(a) // panic: runtime error: index out of range [0] with length 0

	b := make([]int, 4)
	b[0] = 1
	fmt.Println(b)
}

/*
二维切片
*/
func twoDSlice() {
	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + 1
		}
	}
	fmt.Println(twoD) // [[1] [2 2] [3 3 3]]
}
