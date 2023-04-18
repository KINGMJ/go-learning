package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
	fmt.Println("10s后执行...")
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
