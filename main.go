package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()

	fmt.Printf("%05d%d%02d%02d\n", 1, now.Year(), now.Month(), now.Day())

}

// 打印json格式的结构体数据
func prettyJson(req any) {
	data, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling to JSON: %v", err)
	}
	prettyOutput := (string(data))
	fmt.Println(prettyOutput)
}
