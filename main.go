package main

import (
	"fmt"
)

const (
	Secure = 1 >> iota
	Authors
	Reader
	Third
)

func main() {
	fmt.Println(Secure)
	fmt.Println(Authors)
	fmt.Println(Reader)
	fmt.Println(Third)
}
