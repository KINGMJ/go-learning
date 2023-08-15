package main

import "fmt"

func main() {
	demo1()
}

// 编写一个函数，接受一个整数切片作为参数，返回该切片中的最大值和最小值。
func demo1() {
	slice := []int{1, 3, 6, 7, 2, 2, 1}
	max, min := findMaxMin(slice)
	fmt.Printf("最大值：%d，最小值：%d\n", max, min)
}

// 实现指针实现
func findMaxMin(slice []int) (int, int) {
	if len(slice) == 0 {
		return 0, 0
	}
	min := slice[0]
	max := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return max, min
}
