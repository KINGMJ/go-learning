package main

import (
	"fmt"
)

func mapChan(in <-chan any, fn func(any) any) <-chan any {
	out := make(chan any)
	if in == nil {
		close(out)
		return out
	}
	go func() {
		defer close(out)
		for v := range in {
			out <- fn(v)
		}
	}()
	return out
}

func reduce(in <-chan any, fn func(r, v any) any) any {
	if in == nil {
		return nil
	}
	out := <-in
	for v := range in {
		out = fn(out, v)
	}
	return out
}

func asStream(done <-chan any) <-chan any {
	s := make(chan any)
	values := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(s)
		for _, v := range values {
			select {
			case <-done:
				return
			case s <- v:
			}
		}
	}()
	return s
}

func main() {
	in := asStream(nil)
	// map 操作：乘以10
	mapFn := func(v any) any {
		return v.(int) * 10
	}
	reduceFn := func(r, v any) any {
		return r.(int) + v.(int)
	}

	sum := reduce(mapChan(in, mapFn), reduceFn)
	fmt.Println(sum)
}
