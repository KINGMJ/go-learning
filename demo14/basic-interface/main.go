package main

import "fmt"

func main() {
	var dog = new(Dog)
	dog.say()
}

// 定义一个接口
type Sayer interface {
	say()
}

// 定义一个结构体
type Dog struct{}

// 实现接口
func (d Dog) say() {
	fmt.Println("汪汪汪")
}
