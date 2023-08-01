package main

import "fmt"

func main() {
	pointerDemo()
}

/*
指针类型和指针地址
*/
func pointerDemo() {
	v := 5
	// pv 是 v 的指针
	pv := &v
	// pv 的值是 v 在内存中的地址
	fmt.Println(pv)        // 0xc00007c010
	fmt.Printf("%T\n", pv) // *int
}

/*
指针取值
*/
func pointerDemo2() {
	v := 5
	pv := &v
	// 指针取值，根据指针去内存取值
	c := *pv
	fmt.Printf("%T\n", c) // int
	fmt.Printf("%v\n", c) // 5
}

/*
通过指针改变原有的值
*/
func pointerDemo3() {
	v := 5
	pv := &v
	// 通过 pv 改变 v 的值
	*pv = 10
	fmt.Println(v) // 10

}

/*
指针传值
*/
func pointerDemo4() {
	var modify1 = func(x int) {
		x = 100
	}
	// 传递进来的 x 是一个指针，*int 代表 int 的指针类型
	var modify2 = func(x *int) {
		// 对指针进行取值操作
		*x = 100
	}
	a := 10
	modify1(a)
	fmt.Println(a) // 10

	modify2(&a)
	fmt.Println(a) // 100
}

/*
空指针：指针的零值为 nil
*/
func pointerDemo5() {
	var p *string
	fmt.Println(p == nil) // true
}

func pointerDemo6() {
	a := new(int)
	fmt.Printf("%T\n", a) // *int
	fmt.Println(*a)       // 0
}
