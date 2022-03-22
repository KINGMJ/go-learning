package main

import "fmt"

func main() {
	emptyInterface()
}

/*
	空接口
*/
func emptyInterface() {
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
