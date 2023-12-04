package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// 创建一个包含换行符的字符串读取器
	reader := strings.NewReader("Apple\nOrange\nBanana")

	// 创建一个 Scanner，并使用默认的 Split 方法（按行分割）
	scanner := bufio.NewScanner(reader)

	// 使用 Scan 方法逐行扫描数据
	for scanner.Scan() {
		// 获取扫描到的数据
		line := scanner.Text()
		// 对数据进行处理，这里简单打印
		fmt.Println(line)
	}
	// 检查是否有扫描过程中的错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
