package main

import (
	"fmt"
)

func main() {
	interfaceDemo3()
}

// 实现接口
type Sayer interface {
	say()
}
type Dog struct{}

func (d Dog) say() {
	fmt.Println("汪汪汪")
}

func interfaceDemo() {
	var dog = new(Dog)
	dog.say()
}

// 值接收者和指针接收者实现接口的区别
type Mover interface {
	move()
}

func (d Dog) move() {
	fmt.Println("狗会动")
}

// func (d *Dog) move() {
//	fmt.Println("狗会动")
// }

func interfaceDemo2() {
	var x Mover
	var dog1 = Dog{}
	// x 可以接收 Dog 类型
	x = dog1

	// x 可以接收 *Dog 类型
	var dog2 = &Dog{}
	x = dog2
	x.move()
}

// 接口嵌套
type Animal interface {
	Sayer
	Mover
}

type Cat struct {
	name string
}

func (c Cat) say() {
	fmt.Println("喵喵喵")
}
func (c Cat) move() {
	fmt.Println("猫会动")
}

func interfaceDemo3() {
	var x Animal = Cat{name: "花花"}
	x.move()
	x.say()
}

/*
	空接口
*/
func emptyInterfaceDemo() {
	var a interface{}
	var i int = 5
	s := "Hello World!"

	type StructType struct {
		i, j int
		k    string
	}

	a = i
	fmt.Printf("%T\n", a) // int
	//还原
	i, ok := a.(int)
	fmt.Println(ok)        // true
	fmt.Printf("%#v\n", i) // 5

	a = s
	s = a.(string)
	fmt.Println(s) // Hello World!

	a = &StructType{1, 2, "Hello"}
	fmt.Println(a) // &{1 2 Hello}
}
