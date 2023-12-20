package main

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	demo5()
}

func demo1() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		fmt.Println("exec #1")
		return nil
	})

	// 这个任务中又开启了一个协程，在里面通过 ctx.Done 去处理取消操作
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("exec #2")
		go func() {
			select {
			case <-ctx.Done():
				fmt.Println("#2 goroutine canceled")
			case <-time.After(5 * time.Second):
				fmt.Println("#2 goroutine completed")
			}
		}()
		return nil
	})

	// #3 和 #4都返回了一个错误
	g.Go(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("exec #3")
		return errors.New("failed to exec #3")
	})

	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #4")
		return errors.New("failed to exec #4")
	})

	// 调用 wait 等待所有的子任务执行完毕
	if err := g.Wait(); err == nil {
		fmt.Println("Successful exec all")
	} else {
		fmt.Println("failed:", err)
	}
}

func demo2() {
	g := new(errgroup.Group)

	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		// 这里使用 Loop variable capture 处理闭包的变量捕获
		// https://golang.org/doc/faq#closures_and_goroutines
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", body)
			return nil
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successful fetched all URLs.")
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string
type Search func(ctx context.Context, query string) (Result, error)

func fakeSearch(kind string) Search {
	return func(_ context.Context, query string) (Result, error) {
		return Result(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}

func demo3() {
	Google := func(ctx context.Context, query string) ([]Result, error) {
		g, ctx := errgroup.WithContext(ctx)
		searches := []Search{Web, Image, Video}
		results := make([]Result, len(searches))
		for i, search := range searches {
			i, search := i, search
			g.Go(func() error {
				result, err := search(ctx, query)
				if err == nil {
					results[i] = result
				}
				return err
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}
		return results, nil
	}

	results, err := Google(context.Background(), "golang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

func demo4() {
	var g errgroup.Group
	var result = make([]error, 3)

	// 启动第一个子任务,它执行成功
	g.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("exec #1")
		result[0] = nil // 保存成功或者失败的结果
		return nil
	})

	// 启动第二个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(10 * time.Second)
		fmt.Println("exec #2")

		result[1] = errors.New("failed to exec #2") // 保存成功或者失败的结果
		return result[1]
	})

	// 启动第三个子任务，它执行成功
	g.Go(func() error {
		time.Sleep(15 * time.Second)
		fmt.Println("exec #3")
		result[2] = nil // 保存成功或者失败的结果
		return nil
	})

	if err := g.Wait(); err == nil {
		fmt.Printf("Successfully exec all. result: %v\n", result)
	} else {
		fmt.Printf("failed: %v\n", result)
	}
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
func demo5() {
	m, err := MD5All(context.Background(), "../../")
	if err != nil {
		log.Fatal(err)
	}
	for k, sum := range m {
		fmt.Printf("%s:\t%x\n", k, sum)
	}
}

type result struct {
	path string
	sum  [md5.Size]byte
}

func MD5All(ctx context.Context, root string) (map[string][md5.Size]byte, error) {
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan string)

	g.Go(func() error {
		defer close(paths)
		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	})

	c := make(chan result)
	const numDigesters = 20
	for i := 0; i < numDigesters; i++ {
		g.Go(func() error {
			for path := range paths {
				data, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				select {
				case c <- result{path, md5.Sum(data)}:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
	}

	go func() {
		g.Wait()
		close(c)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c {
		m[r.path] = r.sum
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return m, nil
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
