package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		DB:           0,
		Password:     "123456",
		PoolSize:     100,
		MinIdleConns: 10,
	})
	// 心跳检测：使用Ping方法检查连接是否正常
	_, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
}

func main() {
	scene1()
}

// 加锁
func lock(key, value string, expire time.Duration, ctx context.Context) (isGetLock bool, err error) {
	res, err := rdb.Do(ctx, "set", key, value, "px", expire.Milliseconds(), "nx").Result()
	if err != nil {
		return false, err
	}
	if res == "OK" {
		return true, nil
	}
	return false, nil
}

func unlock(key, value string, block time.Duration, ctx context.Context) (isReleased bool, err error) {
	// 监控锁对应的 Key，如果其它的客户端对这个 Key 进行了更改，那么本次事务会被取消。
	err = rdb.Watch(ctx, func(tx *redis.Tx) error {
		// 释放锁之前，校验是否持有锁
		val, err := tx.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			return err
		}
		if val == value {
			// 模拟客户端阻塞n秒，锁超时，自动清除
			time.Sleep(block * time.Second)
			// 客户端恢复，继续释放锁
			_, err := tx.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
				// 删除锁
				pipeliner.Del(ctx, key)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	}, key)

	if err != nil && err != redis.Nil {
		return false, err
	}
	return true, nil
}

// 场景一：3个线程，
// 线程1最开始抢到锁，执行任务，另外两个线程等待
func scene1() {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fmt.Println("#1 开始获取锁：")
		defer wg.Done()
		// 获取分布式锁，过期时间为30s
		isLocked, err := lock("lock_key", "goroutine1", 30000*time.Millisecond, ctx)
		if !isLocked {
			fmt.Printf("#1 获取锁失败，原因：%s", err)
		}
		fmt.Println("#1 开始执行业务...")
		time.Sleep(10 * time.Second)
		// 任务执行完成，解锁
		isReleased, _ := unlock("lock_key", "goroutine1", 3, ctx)
		if isReleased {
			fmt.Println("#1 解锁成功...")
		} else {
			fmt.Println("#1 解锁失败")
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		// fmt.Println("#2 开始获取锁：")
		defer wg.Done()
		// 线程“自旋”，等待锁
		for {
			// 获取分布式锁，过期时间为3s
			isLocked, _ := lock("lock_key", "goroutine2", 3000*time.Millisecond, ctx)
			// 获取到锁，停止"自旋"
			if isLocked {
				fmt.Println("#2 获取到锁")
				break
			}
		}

		fmt.Println("#2 开始执行业务...")
		// 模拟执行任务
		time.Sleep(10 * time.Second)
		// 任务执行完成，解锁
		isReleased, _ := unlock("lock_key", "goroutine2", 3, ctx)
		if isReleased {
			fmt.Println("#2 解锁成功...")
		} else {
			fmt.Println("#2 解锁失败")
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		defer wg.Done()
		// 线程“自旋”，等待锁
		for {
			// 获取分布式锁，过期时间为3s
			isLocked, _ := lock("lock_key", "goroutine3", 3000*time.Millisecond, ctx)
			// 获取到锁，停止"自旋"
			if isLocked {
				fmt.Println("#2 获取到锁")
				break
			}
		}

		fmt.Println("#3 开始执行业务...")
		// 模拟执行任务
		time.Sleep(10 * time.Second)
		// 任务执行完成，解锁
		isReleased, _ := unlock("lock_key", "goroutine3", 3, ctx)
		if isReleased {
			fmt.Println("#3 解锁成功...")
		} else {
			fmt.Println("#3 解锁失败")
		}
	}()

	wg.Wait()
	fmt.Println("所有线程执行完毕...")
}
