package main

import "fmt"

type mapKey struct {
	key int
}

func main() {
	var m = make(map[mapKey]string)
	var key = mapKey{10}
	m[key] = "hello"
	fmt.Printf("m[key]=%s\n", m[key])
	// 修改key的字段的值后再次查询map，无法获取刚才add进去的值
	key.key = 12
	value, ok := m[key]
	fmt.Println(ok) // false
	fmt.Printf("m[key]=%s\n", value)
}
