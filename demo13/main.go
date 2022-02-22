package main

import (
	"fmt"
)

func main() {
	methodsDemo2()
}

/*
	命名返回值
*/
func namedReturnValuedDemo() {
	var add = func(a, b int) (c int) {
		c = a + b
		return
	}
	var a, b int = 1, 2
	c := add(a, b)
	fmt.Println(c)
}

/*
	函数字面量
*/
func functionLiteralsDemo() {
	func(str string) {
		fmt.Println(str)
	}("Hello!")
}

/*
	将函数作为参数传递
*/

func passFunctionAsValueDemo() {
	var add = func(a, b int) int {
		return a + b
	}
	var runFuc = func(a, b int, intOp func(int, int) int) {
		fmt.Printf("intOp(%d, %d) = %d\n", a, b, intOp(a, b))
	}
	runFuc(2, 3, add)
	runFuc(2, 3, func(a, b int) int {
		return a * b
	})
	// intOp(2, 3) = 5
	// intOp(2, 3) = 6
}

/*
	闭包
*/
func closuresDemo() {
	var naturalNumbers = func() func() int {
		i := 0
		f := func() int {
			i++
			return i
		}
		return f
	}

	// n 是一个闭包，它是由一个函数 f 和 一个关联的环境（naturalNumbers 的作用域）组成的
	n := naturalNumbers()
	fmt.Println(n()) // 1
	fmt.Println(n()) // 2

	// 再次调用 naturalNumbers 将创建并返回一个新的函数，这将在 naturalNumbers 中初始化
	// 一个新的 i，这意味着新返回的函数形成了另一个闭包，该闭包具有与函数相同的部分（依然是 f），
	// 但具有全新的环境（新初始化的 i）
	o := naturalNumbers()
	fmt.Println(o()) // 1
	fmt.Println(o()) // 2
}

type User struct {
	Name  string
	Email string
}

func (u User) notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

/*
	方法
*/
func methodsDemo() {
	// 值类型调用
	u1 := User{"go", "golang@golang.com"}
	u1.notify() // go : golang@golang.com

	// 指针类型调用
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.notify() // go : go@go.com
}

type Data struct {
	x int
}

func (self Data) ValueTest() {
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", self)
}

/*
	方法 receiver 为 value 或 pointer 的对比
*/
func methodsDemo2() {
	d := Data{}
	p := &d

	fmt.Printf("Data: %p\n", p)
	d.ValueTest()   // ValueTest(d)
	d.PointerTest() // PointerTest(&d)

	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
	/*
		Pointer 的值都是一样的，Value 的值都不一样，因为是使用的副本

		Data: 0xc000014098
		Value: 0xc0000140a8
		Pointer: 0xc000014098
		Value: 0xc0000140b0
		Pointer: 0xc000014098
	*/
}
