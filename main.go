package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

// Inventory represents the inventory with price as the key and stock as the value
type Inventory map[int]int

// Supplier represents a supplier with a name and the number of items sold
type Supplier struct {
	Name string
	Sold int
}

// CalculateRevenue calculates the revenue for each supplier based on the inventory
func CalculateRevenue(inventory Inventory, suppliers []Supplier) map[string]map[int]int {
	// Sort inventory prices in descending order
	var prices []int
	for price := range inventory {
		prices = append(prices, price)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(prices)))
	prettyJson(prices)

	revenue := make(map[string]map[int]int)

	// Iterate over each supplier and calculate the revenue
	for _, supplier := range suppliers {

		// 供应商的库存
		itemsToSell := supplier.Sold

		// 供应商的售卖
		supplierRevenue := make(map[int]int)

		for _, price := range prices {
			if itemsToSell == 0 {
				break
			}
			// 如果当前价格的库存 > 0
			if inventory[price] > 0 {
				if inventory[price] >= itemsToSell {
					supplierRevenue[price] = itemsToSell
					inventory[price] -= itemsToSell
					itemsToSell = 0
				} else {
					supplierRevenue[price] = inventory[price]
					itemsToSell -= inventory[price]
					inventory[price] = 0
				}
			}
		}
		revenue[supplier.Name] = supplierRevenue
	}
	return revenue
}

func main() {
	// Initial inventory
	inventory := Inventory{

		300: 1,
		100: 1,
		200: 1,
	}

	// Suppliers and their sales
	suppliers := []Supplier{
		{Name: "Supplier A", Sold: 1},
		{Name: "Supplier B", Sold: 2},
	}

	revenue := CalculateRevenue(inventory, suppliers)

	for name, rev := range revenue {
		fmt.Printf("%s revenue: %d\n", name, rev)
	}
}

// 打印json格式的结构体数据
func prettyJson(req any) {
	data, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling to JSON: %v", err)
	}
	prettyOutput := (string(data))
	fmt.Println(prettyOutput)
}
