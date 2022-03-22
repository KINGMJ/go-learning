package main

import (
	"fmt"
	"strconv"
)

func main() {
	float2String()
}

/**
 float 转换为 string
**/
func float2String() {
	// Convert float to string with FormatFloat
	var f32 float32 = 1.3
	bitSize := 32
	s1 := strconv.FormatFloat(float64(f32), 'E', -1, bitSize)
	fmt.Printf("f32: %s\n", s1) // f32: 1.3E+00

	// Convert float to string with Sprintf
	var f64 float64 = 1.54
	s := fmt.Sprintf("%f", f64)
	fmt.Printf("f is: '%s'\n", s) // f is: '1.540000'
}
