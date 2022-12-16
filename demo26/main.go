package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	readFileDemo()
}

func readFileDemo() {
	lines, err := readFileAsLines("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}

func readFileAsLines(path string) ([]string, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := string(d)
	lines := strings.Split(s, "\n")
	return lines, nil
}
