package main

import "fmt"

/*
	值接收者和指针接收者实现接口的区别
*/

func main() {
	var x Mover
	var dog1 = Dog{}
	// 值接收者实现接口，x 可以接收 Dog 类型；指针接收者实现接口，x 不可以接收
	x = dog1

	// 两种方式，x 都可以接收 *Dog 类型
	var dog2 = &Dog{}
	x = dog2
	x.move()
}

type Mover interface {
	move()
}

// 定义一个结构体
type Dog struct{}

// 值接收者
func (d Dog) move() {
	fmt.Println("狗会动")
}

// // 指针接收者
// func (d *Dog) move() {
// 	fmt.Println("狗会动")
// }
