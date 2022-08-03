package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	demo4()
}

func demo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// 练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
func demo2() {
	fmt.Println(os.Args[0])
}

// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
// go run . aa bb cc
func demo3() {
	for index, value := range os.Args {
		if index == 0 {
			continue
		}
		fmt.Printf("index: %v, value: %s\n", index, value)
	}
}

// 练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。
func demo4() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

/*
	1. os.Args的第一个元素：os.Args[0]，是命令本身的名字；其它的元素则是程序启动时传给它的参数
			go run . aa
	2. 字符串连接使用 +
  3. range/continue 的使用
	4. strings.Join 函数替代循环 += ，提升效率
*/
