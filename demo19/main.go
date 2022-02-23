package main

import (
	"fmt"
)

var _ int64 = s()

func main() {
	fmt.Println("function main() --->")
}

func init() {
	fmt.Println("function init() --->")
}

func s() int64 {
	fmt.Println("function s() --->")
	return 1
}
