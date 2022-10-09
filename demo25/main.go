package main

import (
	"fmt"
)

func main() {
	s1 := make([]map[string]interface{}, 0)
	index := 0
	for i := 1; i < 10; i++ {
		m1 := map[string]interface{}{"Foo": 1, "Bar": 30, "Age": 1}
		index++
		m1["Age"] = index
		fmt.Println(m1)
		s1 = append(s1, m1)
	}
	fmt.Println(s1)
}
