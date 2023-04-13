package main

import "fmt"

func main() {
	fmt.Println(gcd(12, 16))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
		fmt.Println(x, y)
	}
	return x
}
