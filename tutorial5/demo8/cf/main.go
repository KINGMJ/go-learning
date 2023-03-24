package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/KINGMJ/go-learning/tutorial5/demo8/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Println(f)
		fmt.Println(c)

		// fmt.Printf("%s = %s, %s = %s\n",
		// 	f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
