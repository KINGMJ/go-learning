package main

import "fmt"

func main() {
	pointerDemo()
}

func pointerDemo() {
	v := 5
	// pv 是 v 的指针
	pv := &v
	fmt.Println(pv) // 0xc00007c010

	// 通过 pv 改变 v 的值
	*pv = 5
	fmt.Println(v)  // 5
	fmt.Println(pv) // 0xc00007c010
}
