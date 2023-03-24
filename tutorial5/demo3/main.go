package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/KINGMJ/go-learning/tutorial5/demo3/gif"
)

const (
	WhiteIndex = iota
	BlackIndex
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// out这个变量是io.Writer类型，这个类型支持把输出结果写到很多目标
	gif.Lissajous(os.Stdout, 5)
}

/**
	1. 运行：go run .>out.gif
**/
