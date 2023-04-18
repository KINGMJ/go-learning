package main

import (
	"fmt"
)

func main() {
	sliceDemo3()
}

func sliceDemo1() {
	slice := make([]int, 12)
	slice = appendInt(slice, 100)
	fmt.Println(slice)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func sliceDemo2() {
	strings := []string{"h", "", "e", "l", "", "l", "o"}
	fmt.Println(strings)
	fmt.Println(noempty2(strings))
}

// 去除数组中的空字符串
//
//	@param strings
//	@return []string
func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func sliceDemo3() {
	stack := []int{1, 2, 3}
	// 压栈
	stack = append(stack, 4)
	// 返回栈顶元素
	top := stack[len(stack)-1]
	fmt.Println(top)
	// 出栈
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
