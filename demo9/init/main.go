package main

import (
	"fmt"
	"time"
)

func main() {
	demo4()
}

func demo1() {
	var m map[int]int
	// 编译期间不会报错，运行时会报错
	m[1] = 2
	fmt.Println(m)
}

// 从一个 nil 的 map对象中获取值不会panic，返回的是一个零值
func demo2() {
	var m map[int]int
	fmt.Println(m[2]) // 0
}

type Counter struct {
	Website      string
	Start        time.Time
	PageCounters map[string]int
}

func NewCounter(website string) *Counter {
	return &Counter{
		Website:      website,
		Start:        time.Now(),
		PageCounters: make(map[string]int),
	}
}

func demo3() {
	var c Counter
	c.Website = "www.baidu.com"
	c.PageCounters["/"]++

	fmt.Println(c)
}

func demo4() {
	c := NewCounter("www.baidu.com")
	fmt.Println(c)
}
