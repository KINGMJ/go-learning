package main

import (
	"fmt"
	"time"

	"github.com/thoas/go-funk"
)

// 定义ProductStock结构体
type ProductStock struct {
	ID        int64     `json:"id"`
	StoreID   int64     `json:"storeID"`
	ProductID int64     `json:"productID"`
	Count     int64     `json:"count"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

func main() {
	// 假设这是你的map数据
	productStocks := map[int64]*ProductStock{
		1203: {
			ID:        1203,
			StoreID:   30,
			ProductID: 123592,
			Count:     454,
			Created:   time.Date(2024, 5, 9, 18, 32, 0, 0, time.Local),
			Updated:   time.Date(2024, 6, 11, 13, 49, 46, 0, time.Local),
		},
		1204: {
			ID:        1204,
			StoreID:   30,
			ProductID: 123593,
			Count:     560,
			Created:   time.Date(2024, 5, 9, 18, 32, 45, 0, time.Local),
			Updated:   time.Date(2024, 6, 5, 16, 27, 42, 0, time.Local),
		},
		1211: {
			ID:        1211,
			StoreID:   30,
			ProductID: 123596,
			Count:     306,
			Created:   time.Date(2024, 5, 11, 16, 42, 43, 0, time.Local),
			Updated:   time.Date(2024, 6, 5, 16, 27, 42, 0, time.Local),
		},
	}

	keys := funk.Keys(productStocks).([]int64)
	fmt.Println(productStocks[keys[0]])
}
