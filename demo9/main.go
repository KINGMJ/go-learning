package main

import "fmt"

func main() {
	sliceAsMapValueDemo()
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

/*
	map 作为 slice 的元素
*/
func mapAsSliceItemDemo() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		// index: 0 value: map[]
		fmt.Printf("index: %d value: %v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的 map 元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		// index: 0 value: map[address:红旗大街 name:王五 password:123456]
		fmt.Printf("index: %d value: %v\n", index, value)
	}
}

/*
	slice 作为 map 的值
*/
func sliceAsMapValueDemo() {
	// map 的 value 是一个 []string slice
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap) // map[]
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap) // map[中国:[北京 上海]]
}
