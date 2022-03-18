package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 函数首字母大写
// go 函数可以返回多个值
func Hello(name string) (string, error) {
	if name == "" {
		// 返回错误信息
		return "", errors.New("empty name")
	}
	// := 运算符是在一行中声明和初始化变量，根据右边返回的返回值来确定类型
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	// 添加一个 nil 作为成功返回的第二个值，以满足函数的返回值要求
	return message, nil
}

// 参数是一个slice，返回的格式是一个 map 和 error
func Hellos(names []string) (map[string]string, error) {
	// 初始化一个map，格式：make(map[key-type]value-type).
	messages := make(map[string]string)
	// 迭代，index，value。由于不需要index，可以使用 The blank identifier 替代
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init 函数，go 程序启动时，在全局变量被初始化之后，自动执行 init 函数
func init() {
	rand.Seed(time.Now().UnixNano())
}

// 小写字母开头的方法只能在包内部被使用
func randomFormat() string {
	//go 中使用切片（Slice）表示一个动态数组
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

/**
 1. go 的方法可以返回多个值，为了保证返回值一致，可以使用 nil 关键字。比如： messages, nil
 2. 对外暴露的方法，方法名首字母大写；对内的首字母小写
 3. 初始化一个map，格式：make(map[key-type]value-type).
 4. 使用 for 对 map 进行迭代，参数为 index, value ，如果不需要index，可以使用 _ (The blank identifier)
**/
