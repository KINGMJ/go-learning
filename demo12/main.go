package main

import (
	"fmt"
)

func main() {
	rangeStringDemo()
}

func rangeStringDemo() {
	s := "Hey 世界"
	for i, r := range s {
		fmt.Printf("idx: %d, rune: %d\n", i, r)
	}
}
