// declare a main package
package main

// 导入 fmt 包
import (
	"fmt"
	"log"
	// go mod tidy 安装一个外部包
	"rsc.io/quote"
	// 自己写的包，本地使用 go mod edit -replace go-learning/greetings=../greetings
	"go-learning/greetings"
)

// mian 方法是默认执行的方法，使用 go run . 运行
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// 使用日志
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// 调用另一个模块的方法
	// go 常用错误处理，将错误作为值返回，以便调用者可以检查它

	// message, err := greetings.Hello("")
	message, err := greetings.Hello("Gladys")
	if err != nil {
		// 打印错误并停止程序
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}

/**
 1. 运行： go run .
 2. 安装模块：go mod tidy
 3. 使用本地的模块替代生产环境的：go mod edit -replace go-learning/greetings=../greetings
 4. := 运算符是在一行中声明和初始化变量，根据右边返回的返回值来确定类型
 5. 插值运算符 fmt.Sprintf("Hi, %v. Welcome!", name)
 6. 错误处理，将错误作为返回值返回
 7. init 方法
 8. slice 动态数组
**/
