package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	structDemo6()
}

func typeDemo() {
	type myInt int
	var a myInt = 12
	fmt.Printf("%T\n", a) // main.myInt

	type myByte = uint8
	var b myByte = 1
	fmt.Printf("%T\n", b) // uint8
}

func structsDemo() {
	type Person struct {
		name string
		city string
		age  int8
	}

	// 使用 var 实例化结构体
	var person Person
	person.name = "Jack"
	person.city = "Shanghai"
	person.age = 12
	fmt.Println(person) // {Jack Shanghai 12}

	// 使用字面量方式，省略写法必须保证顺序一致，并且要列出全部字段
	person1 := Person{name: "Mary", age: 12}
	person2 := Person{"Rose", "New York", 13}

	fmt.Println(person1)      // {Mary  12}
	fmt.Println(person1.city) // ""
	fmt.Println(person2)      // {Rose New York 13}

	// 使用 new 得到的是 Person 类型的零值
	var person3 = new(Person)
	person3.name = "Bill"
	person3.city = "Houston"
	fmt.Println(person3)      // &{Bill Houston 0}
	fmt.Println(person3.city) // Houston
	// 可以对指针取值，改变原始值
	*person3 = Person{"Lily", "Beijing", 12}
	fmt.Println(person3) // &{Lily Beijing 12}

	// 使用 & 取地址操作
	var person4 = &Person{}
	person4.name = "Jason"
	fmt.Println(person4) // &{Jason  0}
}

func structsDemo2() {
	var person struct {
		name string
		city string
		age  int
	}
	person.name = "Bill"
	fmt.Println(person)

	// 字面量方式
	person1 := struct {
		name string
		city string
		age  int
	}{"Bill", "Houston", 8}

	fmt.Println(person1)
}

// 构造函数demo

type Person struct {
	name string
	age  int8
}

// 模拟构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func structDemo3() {
	p1 := NewPerson("jack", 25)
	p1.Dream()
}

// 嵌套结构体

type Request struct {
	Resource string
}

type AuthenticatedRequest struct {
	Request            Request
	Username, Password string
}

func structDemo4() {
	ar := new(AuthenticatedRequest)
	ar.Request.Resource = "example.com/requesr"
	ar.Username = "bob"
	ar.Password = "Pass"
	fmt.Printf("%#v\n", ar)
}

// 嵌套匿名结构体
type Address struct {
	Province, City string
}

type User struct {
	Name, Gender string
	Address      // 匿名结构体
	City         int
}

func structDemo5() {
	var user User
	user.Name = "Jack"
	user.Gender = "femail"
	user.Province = "湖北"
	//当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找
	user.City = 001
	user.Address.City = "武汉"
}

// 结构体标签

type Account struct {
	Username      string `json:"username"`
	DisplayName   string `json:"display_name"`
	FavoriteColor string `json:"favorite_color,omitempty"`
}

func structDemo6() {
	a1 := Account{
		Username:      "jack li",
		DisplayName:   "JL",
		FavoriteColor: "green",
	}
	data, err := json.Marshal(a1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) // json str:{"username":"jack li","display_name":"JL","favorite_color":"green"}
}
