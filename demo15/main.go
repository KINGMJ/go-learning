package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	errorDemo3()
}

// 错误处理
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("invalid argument '%f', must be >= 0", n)
	}
	return math.Sqrt(n), nil
}

func printSqrt(n float64) {
	if res, err := sqrt(n); err == nil {
		fmt.Printf("sqrt of %f is %f\n", n, res)
	} else {
		fmt.Printf("sqrt of %f returned error '%s'\n", n, err)
	}
}

func errorDemo() {
	printSqrt(-1)
	printSqrt(1)
}

// 自定义错误类型
type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func customErrorDemo() {
	var err error = (*MyError)(nil)
	fmt.Printf("errorType: %T, errValue: %s\n", err, err)
}

// 自定义错误类型的陷阱：不要返回基础类型，始终返回 error 类型
type CustomError struct {
}

func (e *CustomError) Error() string {
	return "this is a custom error"
}

func errorDemo2() {
	var err1 error
	var err2 error = nil
	var err3 error = (*CustomError)(nil)
	fmt.Println(err1 == nil && err2 == nil) // true
	fmt.Printf("errorType: %T, errValue: %s\n", err3, err3)
	// errorType: *main.CustomError, errValue: this is a custom error
}

// errors.New

func errorDemo3() {
	var errorType1 = errors.New("EOF")
	var errorType2 = errors.New("EOF")
	fmt.Println(errorType1 == errorType2) // false
}
