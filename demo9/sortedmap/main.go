package main

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

func main() {
	demo2()
}

func demo1() {
	m := map[string]string{
		"A": "jack", "B": "rose", "C": "mike",
	}

	for i := 0; i < 100; i++ {
		for key, name := range m {
			fmt.Printf("%s: %s ", key, name)
		}
		fmt.Println("")
	}
}

func sortedMap[K constraints.Ordered, V any](m map[K]V, f func(k K, v V)) {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, k := range keys {
		f(k, m[k])
	}
}

func demo2() {
	m := map[string]string{
		"A": "jack", "B": "rose", "C": "mike",
	}
	for i := 0; i < 100; i++ {
		sortedMap(m, func(k string, v string) {
			fmt.Printf("%s: %s ", k, v)
		})
		fmt.Println("")
	}
}
