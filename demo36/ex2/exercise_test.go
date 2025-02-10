package ex2

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	c := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increament()
		}()
	}
	wg.Wait()
	fmt.Println(c.Get())
}

func TestCache(t *testing.T) {
	c := NewCache()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.Set("key1", "value1", time.Second*2)
		c.Set("key2", "value2", time.Second*3)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		c.Set("key3", "value3", time.Second*3)
		v, ok := c.Get("key1")
		fmt.Println(v, ok)
	}()

	// 10s 后获取map
	time.Sleep(time.Second * 10)
	fmt.Printf("%#v\n", c.Data())
	wg.Wait()
}
