package main

import "fmt"

func main() {
	copyMapDemo()
}

func mapDemo() {
	var m1 map[int]int
	fmt.Printf("%#v\n", m1) // map[int]int(nil)
	// m1[1] = 2  // 报错
	fmt.Println(m1[1]) // 0

	// map 声明后是一个 nil 的零值，必须要使用 make 初始化
	if m1 == nil {
		fmt.Println("map is nil. Going to make one.")
		m1 = make(map[int]int)
	}
	fmt.Printf("%#v\n", m1) // map[int]int{}
}

/*
	获取 map 的值
*/
func getMapValueDemo() {
	m := map[string]string{"foo": "foo_value", "bar": ""}
	fmt.Println(m["bar"]) // ""
	// 不存在的键
	fmt.Println(m["not exists"]) //""

	_, hasKey := m["not exists"]
	fmt.Println(hasKey) // false
}

/*
	map 的深拷贝与浅拷贝
*/
func copyMapDemo() {
	src := map[string]int{"one": 1, "two": 2}
	// 浅拷贝
	dst := src
	// 深拷贝
	dst1 := map[string]int{}
	for key, value := range src {
		dst1[key] = value
	}
	src["three"] = 3

	fmt.Println(dst)  // map[one:1 three:3 two:2]
	fmt.Println(dst1) // map[one:1 two:2]
}
