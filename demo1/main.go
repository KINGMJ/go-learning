package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {
	booleanDemo()
}

/*
 布尔类型
*/
func booleanDemo() {
	// bool 类型，true or false
	var b bool = true
	fmt.Println(b)
	// zero value 是 false
	var b2 bool
	fmt.Println(b2) // false

	// bool 的 size 是 1
	fmt.Println(unsafe.Sizeof(b2)) // 1
}

/*
 整数类型
*/
func integerDemo() {
	// int 类型
	var i1 int = 12
	fmt.Println(i1)

	// 打印数据类型
	fmt.Println(reflect.TypeOf(i1)) // int

	// zero value 是 0
	var i2 int
	fmt.Println(i2) // 0
}

/*
 整数转字符串
*/
func integer2StringDemo() {
	// 类型转换
	// Convert int to string with strconv.Itoa
	var i3 int = -38
	var s3 = strconv.Itoa(i3)
	fmt.Println(s3)                 // -38
	fmt.Println(reflect.TypeOf(s3)) // string

	// Convert int to string with fmt.Sprintf
	i4 := fmt.Sprintf("%d", i3)
	fmt.Println(i4)
}

/*
 字符串转整数
*/
func string2IntegerDemo() {
	// Convert string to int with strconv.Atoi
	// s5 := "-38a" // 这种格式转换失败
	s5 := "-38"
	i5, error := strconv.Atoi(s5)
	if error != nil {
		log.Fatalf("strconv.Atoi() failed with %s\n", error)
	}
	fmt.Printf("i5: %d\n", i5) // i5: -38

	// Convert string to int with fmt.Sscanf
	s6 := "345a"
	var i6 int
	_, err := fmt.Sscanf(s6, "%d", &i6)
	if err != nil {
		log.Fatalf("fmt.Sscanf() failed with '%s'\n", err)
	}
	fmt.Printf("i6: %d\n", i6) // 345
}

/**
 float 转换为 string
**/
func float2StringDemo() {
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

/*
 string 转 float
*/
func string2FloatDemo() {
	// Convert string to float with ParseFloat

	// s := "1.234f" // 不支持，报错
	s := "1.23418881" // 1.234189 精度为 6为，向上取整
	f64, error := strconv.ParseFloat(s, 64)
	if error != nil {
		log.Fatalf("strconv.ParseFloat() failed with '%s'\n", error)
	}
	fmt.Printf("f64: %f\n", f64)

	// Convert string to float with Sscanf
	s1 := "1.234f"
	var f float64
	_, err := fmt.Sscanf(s1, "%f", &f)
	if err != nil {
		log.Fatalf("fmt.Sscanf failed with '%s'\n", err)
	}
	fmt.Printf("f: %f\n", f) // 1.234000
}
