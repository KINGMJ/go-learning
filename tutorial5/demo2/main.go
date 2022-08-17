package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	demo2()
}

// 从标准输入中读取数据，检查重复行
func demo1() {
	// 声明一个 map
	counts := make(map[string]int)
	// 扫描 Stdin，它读取输入并将其拆成行或单词；通常是处理行形式的输入最简单的方法。
	// 该变量从程序的标准输入中读取内容
	input := bufio.NewScanner(os.Stdin)
	// 每次调用input.Scan()，即读入下一行，并移除行末的换行符
	// Scan函数在读到一行时返回true
	for input.Scan() {
		// 读取的内容可以调用input.Text()得到
		// 控制循环退出：输入 end 退出循环，否则一直监听输入
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			// \t 制表符；\n 换行符
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 统计文件中的重复行
// go run . a.txt b.txt
func demo2() {
	counts := make(map[string]int)
	// 获取参数，可能传递多个文件
	files := os.Args[1:]
	// 如果没有传递文件，从标准输入中统计；否则从文件中统计
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open函数返回两个值。第一个值是被打开的文件(*os.File），其后被Scanner读取。
			// 第二个值是内置 error 类型的值
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// map作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），
// 被调用函数对map底层数据结构的任何修改，调用者函数都可以通过持有的map引用看到。
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

/*
	1. map 的使用
	2. bufio 包
*/
