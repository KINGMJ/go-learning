package main

import (
	"fmt"
)

/*
	接口嵌套
*/

func main() {
	var x Animal = Cat{name: "花花"}
	x.move()
	x.say()
}

type Mover interface {
	move()
}

type Sayer interface {
	say()
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
