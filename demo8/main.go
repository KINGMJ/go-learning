package main

import "fmt"

func main() {
	nilSliceDemo()
}

func sliceDemo() {
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
	创建一个 slice
*/
func sliceDemo2() {
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

func nilSliceDemo() {
	s := []int(nil)
	if len(s) == 0 {
		fmt.Printf("s  is empty: %#v\n", s)
	}

	var s2 []int
	if len(s2) == 0 {
		fmt.Printf("s2 is empty: %#v\n", s2)
	}

	s3 := make([]int, 0)
	if len(s2) == 0 {
		fmt.Printf("s3 is empty: %#v\n", s3)
	}
}

/*
	创建静态预分配切片
*/
func sliceDemo3() {
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
func sliceDemo4() {
	a := make([]int, 0, 12)
	fmt.Printf("a has lenght %d and capacity %d\n", len(a), cap(a)) // a has lenght 0 and capacity 12
	for i := 0; i < 16; i++ {
		a = append(a, i)
	}
	fmt.Printf("a has lenght %d and capacity %d\n", len(a), cap(a)) // a has lenght 16 and capacity 24
}

/*
	切片过滤
*/
func filterSliceDemo() {
	var filterEven = func(a []int) []int {
		var res []int
		for _, el := range a {
			if el%2 == 0 {
				continue
			}
			res = append(res, el)
		}
		return res
	}
	a := []int{1, 2, 3, 4, 5}
	res := filterEven(a)
	fmt.Printf("%#v\n", res) // []int{1, 3, 5}
}

/*
  slice 基本表达式
*/
func sliceOperation() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]

	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(s) // [2 3 4]
	fmt.Println(cap(s))
}

/*
	slice 完整表达式
*/
func sliceOperation2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:4:4]

	fmt.Println(a)      // [1 2 3 4 5]
	fmt.Println(s)      // [1 2 3 4]
	fmt.Println(cap(s)) // 4
}

/*
	切片表达式会影响切片的底层数组
*/
func sliceOperation3() {
	// a、s1、s2 底层都是共用同一个数组
	var a [10]int
	// s1 取的是[3,7)
	s1 := a[3:7] // [0 0 0 [ 0 0 0 0 ] 0 0 0]
	// s2 是在 s1 基础上取的[1,4)
	s2 := s1[1:4] // [0 0 0 [ 0 [0 0 0] ] 0 0 0]

	// 对 s2 进行修改，实际操作的是底层数组，都会受到影响
	s2[1] = 42

	fmt.Println(s2) // [0 42 0]
	fmt.Println(s1) // [0 0 42 0]
	fmt.Println(a)  // [0 0 0 0 0 42 0 0 0 0]
}

/*
	append a slice to slice
*/
func appendSliceDemo() {
	a := []string{"!"}
	a2 := []string{"Hello", "World"}
	a = append(a, a2...)
	fmt.Printf("a: %#v\n", a) // a: []string{"!", "Hello", "World"}

	// 改变了 a2，不会影响 a
	a2[1] = "Jack"
	fmt.Println(a)  // [! Hello World]
	fmt.Println(a2) // [Hello Jack]

}

/*
	copy 操作
*/
func copySliceDemo() {
	// 原 slice
	slice := []int{1, 2, 3, 4, 5}
	// 通过切片表达式生成的 slice1，所以 slice1 与 slice 内部共享一个数组
	slice1 := slice[:]

	slice2 := []int{5, 4, 3}

	// 把 slice2 copy 给 slice，只会复制 slice2 的前三个元素到 slice 的前三个元素
	copy(slice, slice2)

	fmt.Println(slice2) // [5 4 3]
	fmt.Println(slice)  // [5 4 3 4 5]
	fmt.Println(slice1) // [5 4 3 4 5]

	// slice2 的修改不会影响 slice
	slice2[0] = 6
	slice1[0] = 10

	fmt.Println(slice2) // [6 4 3]
	fmt.Println(slice)  // [10 4 3 4 5]
	fmt.Println(slice1) // [10 4 3 4 5]
}

/*
	使用 append 到一个空 slice 中替代 copy
*/
func copySliceDemo2() {
	src := []int{1, 2, 3}
	dst := append([]int{}, src...)

	// 改变原 slice，不受影响
	src[0] = 4
	fmt.Println(src) // [4 2 3]
	fmt.Println(dst) // [1 2 3]
}

/*
	删除开头的位置
*/
func removeSliceDemo() {
	s := []int{10, 11, 12, 13}
	n := 2
	// 删除开头的 n 个元素
	//s = s[n:]
	// s = append(s[:0], s[n:]...)

	// copy 返回的是复制的元素数量，也就是 n
	len := copy(s, s[n:])
	fmt.Println(len) // 2
	fmt.Println(s)   // [12 13 12 13]
	// 取[0,2)
	s = s[:len]
	fmt.Println(s) // [12 13]

}

/*
	删除中间的位置
*/
func removeSliceDemo2() {
	s := []int{10, 11, 12, 13}

	i := 1
	n := 2

	// s = append(s[:i], s[i+n:]...)
	len := copy(s[i:], s[i+n:]) // 1，s：[10 13 12 13]
	s = s[:i+len]
	fmt.Println(s) // [10 13]
}

/*
	使用优化地方式移除元素
*/
func removeSliceDemo3() {
	s := []int{10, 11, 12, 13}
	i := 1 // 索引

	lastIdx := len(s) - 1 // 3
	s[i] = s[lastIdx]     // [10 13 12 13]
	s = s[:lastIdx]       // [10 13 12]
	fmt.Printf("s: %#v\n", s)
}
