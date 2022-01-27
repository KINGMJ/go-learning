package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	stringLenDemo()
}

/*
 使用 byte 迭代 string
*/
func byteDemo() {
	s := "str"
	for i := 0; i < len(s); i++ {
		c := s[i]
		fmt.Printf("Byte at index %d is '%c' (0x%x)\n", i, c, c)
	}
}

/*
 使用 runes 迭代 string
*/
func runesDemo() {
	s := "日本語"
	for i, runeChar := range s {
		fmt.Printf("Rune at byte position %d is %#U\n", i, runeChar)
	}
}

func stringDemo() {
	var s string
	s1 := "string\nliteral\nwith\tescape characters\n"
	s2 := `raw string literal
which doesnt't recgonize escape characters like \n
`
	fmt.Println(s) // “”
	fmt.Printf("sum of strings\n'%s'\n", s+s1+s2)
}

func stringLenDemo() {
	s := "您好a"
	// len 返回的字符的字节长度，中文在 UTF-8 中占 3 个字节
	fmt.Println(len(s)) // 7
	// 返回字符的长度
	fmt.Println(utf8.RuneCountInString(s))
}
