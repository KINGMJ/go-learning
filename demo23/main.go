/*
 * @Author: KINGMJ 328047478@qq.com
 * @Date: 2022-09-16 17:41:55
 * @LastEditors: KINGMJ 328047478@qq.com
 * @LastEditTime: 2022-09-20 10:32:35
 * @FilePath: /go-learning/demo23/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

type Res struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	structToJsonDemo()
}

// json 转 struct
func jsonToStructDemo() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Res{}
	json.Unmarshal([]byte(str), &res)
	fmt.Printf("%#v\n", res)
	fmt.Println(res.Fruits[0])
}

// struct 转 json
func structToJsonDemo() {
	user := &User{Name: "Frank", Age: 12, Sex: 1}
	b, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
