package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"
)

func main() {
	stringLength()
}

/*
	字符串的长度
*/
func stringLength() {
	// len 获取的是字节长度
	fmt.Println(len("你好"))  // 6
	fmt.Println(len("abc")) // 3

	// 字符的长度和编码有关，使用编码库来获取
	fmt.Println(utf8.RuneCountInString("你好"))    // 2
	fmt.Println(utf8.RuneCountInString("你好abc")) //5

}

/*
 字符串转整数
*/
func string2Integer() {
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

/*
 string 转 float
*/
func string2Float() {
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
