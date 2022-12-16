package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	scanTextDemo3()
}

/*
index 方法
*/
func indexDemo() {
	s := "where hello is ?"
	toFind := "hello"
	idx := strings.Index(s, toFind)
	fmt.Printf("'%s' is in s starting at position %d\n", toFind, idx)

	// 没有找到，返回 -1
	idx = strings.Index(s, "not present")
	fmt.Println(idx)
}

/*
last index 方法
*/
func lastIndexDemo() {
	s := "hello and second hello"
	toFind := "hello"
	idx := strings.LastIndex(s, toFind)
	fmt.Printf("when searching from end, '%s' is in s at position %d\n", toFind, idx)
}

/*
找到所有的 index
*/
func findAllIndexDemo() {
	s := "first is, second is, third is"
	toFind := "is"
	currStart := 0
	for {
		idx := strings.Index(s, toFind)
		if idx == -1 {
			break
		}
		fmt.Printf("found '%s' at position %d\n", toFind, currStart+idx)
		currStart += idx + len(toFind)
		fmt.Println(currStart)
		// 类似于 substring
		s = s[idx+len(toFind):]
		fmt.Println(s)
	}
}

func compareDemo() {
	// 返回 0，-1，1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "A")) // 1

	// 使用 ==，> 和 < 比较更快
	fmt.Println("a" > "A") //1

	// 不区分大小写的比较
	fmt.Println(strings.EqualFold("a", "A")) // true
}

/*
大小写转换
*/
func lowerUpperDemo() {
	//s := "mixed Case"
	//s := "хлEб"
	s := "ǳ"
	fmt.Println(strings.ToLower(s)) // mixed case | хлeб  |  ǳ
	fmt.Println(strings.ToUpper(s)) // MIXED CASE | ХЛEБ  |  Ǳ
	fmt.Println(strings.ToTitle(s)) // MIXED CASE | ХЛEБ  |  ǲ
	fmt.Println(strings.Title(s))   // Mixed Case | ХлEб  |  ǲ
}

func toIntegerDemo() {
	v := "-10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int, -10
	}

	if s, err := strconv.ParseInt(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) //int64, -10
	}
}

func toFloatDemo() {
	v := "3.1415926535"
	if s, err := strconv.ParseFloat(v, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 3.1415927410125732
	}
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 3.1415926535
	}
	if s, err := strconv.ParseFloat("NaN", 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, NaN
	}
}

func trimDemo() {
	s := " str \n "
	s1 := "abacdda"
	s2 := "¡¡¡Hello, Gophers!!!"
	// 删除字符串前后的空格
	fmt.Println(strings.TrimSpace(s)) // str
	// 删除字符串前的指定字符，必须要匹配
	fmt.Println(strings.TrimPrefix(s1, "aba")) // cdda
	// 删除字符串后的指定字符，必须要匹配
	fmt.Println(strings.TrimSuffix(s1, "da")) // abacd
	// 删除字符串中所有的指定字符，不用完全匹配
	fmt.Println(strings.Trim(s2, "¡!a")) // Hello, Gophers
	// 删除字符串左侧所有的指定字符，不用完全匹配
	fmt.Println(strings.TrimLeft("!¡¡¡Hello!, Gophers!!!", "!¡")) // Hello!, Gophers!!!
	// 删除字符串右侧所有的指定字符，不用完全匹配
	fmt.Println(strings.TrimRight("!¡¡¡Hello!, Gophers!!!", "!¡")) // !¡¡¡Hello!, Gophers
}

func replaceDemo() {
	// n = 2，替换2次
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) // oinky oinky oink

	// n < 0，替换次数没有限制，也就是全替换
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) // moo moo moo
	// 等价于下面
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
}

func splitAndJoinDemo() {
	s := "this is a string"
	a := strings.Split(s, " ")
	fmt.Printf("a: %#v\n", a) // a: []string{"this", "is", "a", "string"}

	s2 := strings.Join(a, ",")
	fmt.Printf("s2: %#v\n", s2) // s2: "this,is,a,string"
}

func formatTextDemo() {
	var i int = 23
	fmt.Printf("%v\n", i) // 23
	fmt.Printf("%T\n", i) // int

	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

	fmt.Printf("%T\n", person) // struct { Name string; Age int }
}

// Scan 扫描从标准输入读取的文本，将连续的空格分隔值存储到连续的参数中。
// 换行算作空格。它返回成功扫描的项目数。如果它小于参数的数量，err 将报告原因。
func scanTextDemo() {
	s := "48 123.45s"
	var f float64
	var i int
	nParsed, err := fmt.Sscanf(s, "%d %f", &i, &f)
	if err != nil {
		log.Fatalf("first fmt.Sscanf failed with %s\n", err)
	}
	fmt.Printf("i: %d, f: %f, extracted %d values\n", i, f, nParsed)
}

func scanTextDemo2() {
	s := "48 123.45 56 s"
	var a, b int
	var c float64
	count, err := fmt.Sscan(s, &a, &c, &b)
	if err != nil {
		log.Fatalf("first fmt.Sscanf failed with %s\n", err)
	}
	fmt.Println("数量", count)
	fmt.Printf("解析后的数据: %d, %d, %f", a, b, c)
}

func scanTextDemo3() {
	s := "48 123.45 56 s"
	var a, b int
	var c float64
	count, err := fmt.Sscanln(s, &a, &c, &b)
	if err != nil {
		log.Fatalf("first fmt.Sscanf failed with %s\n", err)
	}
	fmt.Println("数量", count)
	fmt.Printf("解析后的数据: %d, %d, %f", a, b, c)
}
