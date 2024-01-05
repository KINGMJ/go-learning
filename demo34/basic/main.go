package main

import (
	"fmt"
	"reflect"
)

func main() {
	demo1()
}

type ExampleStruct struct {
	Name string
	Age  int
}

func demo1() {
	// example := ExampleStruct{"John", 25}
	var example float64 = 3.2
	// example := make([]int, 12)
	// 获取 reflex.Type 对象
	t := reflect.TypeOf(example)
	v := reflect.ValueOf(example)
	// 获取底层基础类型
	kind := t.Kind()
	fmt.Printf("Value: %v\n", v)
	fmt.Printf("Type: %v, Kind: %v, Name: %v, Size: %v\n", t, kind, t.Name(), t.Size())

	fmt.Println(v.Float())
}

func demo2() {
	var x float64 = 3.4
	// 反射认为下面是指针类型，不是float类型
	reflectSetValue(&x)
	fmt.Println("main:", x)
}

func reflectSetValue(a any) {
	v := reflect.ValueOf(a)
	k := v.Kind()

	switch k {
	case reflect.Float64:
		// 反射修改值
		if v.CanSet() {
			v.SetFloat(6.9)
			fmt.Println("a is ", v.Float())
		} else {
			fmt.Println("a cannot be set.")
		}

	case reflect.Ptr:
		// Elem() 获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println("pointer is :", v.Pointer())
	}
}

func demo3() {
	// var a User = User{Id: 1, Name: "jack", Age: 12}
	m := Boy{User{1, "jack", 20}, "127.0.0.1"}
	// var a float64 = 3.14
	Poni(m)
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello")
}

func (u User) Hellome(name string, age int) {
	fmt.Printf("Hello, I'm %s, I am %d\n", name, age)
}

func Poni(o any) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	fmt.Println("基础类型：", t.Kind())
	fmt.Println("值：", v)
	fmt.Println("字段个数：", t.NumField())
	fmt.Println("函数个数", t.NumMethod())

	// 获取结构体的每个字段
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s:%v，", f.Name, f.Type)
		// 获取字段的值信息
		val := v.Field(i).Interface()
		fmt.Println("Value:", val)
	}
	fmt.Println("=================方法====================")

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("Method name:", m.Name)
		fmt.Println("Method type:", m.Type)
	}
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

func demo4() {
	m := Boy{User{1, "jack", 20}, "127.0.0.1"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}

func demo5() {
	m := Boy{User{1, "jack", 20}, "127.0.0.1"}
	userValue := reflect.ValueOf(&m).Elem().FieldByName("User")

	nameVlaue := userValue.FieldByName("Name")
	if nameVlaue.Kind() == reflect.String {
		nameVlaue.SetString("rose")
	}
	fmt.Println(m)
}

func demo6() {
	u := User{1, "5lmh.com", 20}
	v := reflect.ValueOf(u)
	// 获取方法
	m := v.MethodByName("Hello")
	args := []reflect.Value{}
	m.Call(args)

	m1 := v.MethodByName("Hellome")
	args1 := []reflect.Value{reflect.ValueOf("rose"), reflect.ValueOf(12)}
	m1.Call(args1)
}

type Student struct {
	Name string `json:"name" db:"name"`
}

// var s Student
// v := reflect.ValueOf(&s)
// // 类型
// t := v.Type()
// // 获取字段
// f := t.Elem().Field(0)
// fmt.Println(f.Tag.Get("json"))
// fmt.Println(f.Tag.Get("db"))

func demo7() {
	var s Student = Student{"jack"}
	v := reflect.ValueOf(&s)

	t := v.Type()
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}
