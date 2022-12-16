package main

import (
	"fmt"
)

func main() {
	v := 5
	pv := &v
	// 指针取值，根据指针去内存取值
	c := *pv
	c = 10
	fmt.Printf("%v\n", c) // 5
	fmt.Printf("%v\n", v) // 5
}
