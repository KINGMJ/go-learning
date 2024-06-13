package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/thoas/go-funk"
)

func main() {
	demo2()
}

func demo2() {
	// 定义两个集合
	setA := []int{1, 2, 3, 4}
	setB := []int{1, 2, 2}

	// 计算差集
	diffA, diffB := funk.DifferenceInt(setA, setB)

	// 打印结果
	fmt.Println(diffA, diffB) // 输出: [3 4]
}

type SupplierTotalSales struct {
	SupplierId    int64
	SalesQuantity int
}

type SupplierProductStockLog struct {
	Id             int64 `json:"id"`
	ProductId      int64 `json:"productId"`
	SupplierId     int64 `json:"supplierId"`
	StoreId        int64 `json:"storeId"`
	Stock          int   `json:"stock"`
	SupplierMode   int8  `json:"supplierMode"`
	DeductionStock int   `json:"deductionStock"`
	IsLastDeducted int8  `json:"isLastDeducted"`
}

func demo1() {
	jsonData := `
	[
  {
    "id": 80,
    "productId": 123593,
    "supplierId": 14,
    "supplierMode": 1,
    "storeId": 30,
    "stock": 9,
    "deductionStock": 1,
    "isLastDeducted": 0,
    "date": "2024-05-29T00:00:00+08:00",
    "created": "2024-05-29T10:06:52+08:00",
    "updated": "2024-05-29T10:06:52+08:00"
  },
  {
    "id": 81,
    "productId": 123595,
    "supplierId": 14,
    "supplierMode": 1,
    "storeId": 30,
    "stock": 11,
    "deductionStock": 1,
    "isLastDeducted": 0,
    "date": "2024-05-29T00:00:00+08:00",
    "created": "2024-05-29T10:06:52+08:00",
    "updated": "2024-05-29T10:06:52+08:00"
  },
  {
    "id": 81,
    "productId": 123593,
    "supplierId": 15,
    "supplierMode": 2,
    "storeId": 30,
    "stock": 4,
    "deductionStock": 1,
    "isLastDeducted": 1,
    "date": "2024-05-29T00:00:00+08:00",
    "created": "2024-05-29T10:06:52+08:00",
    "updated": "2024-05-29T10:06:52+08:00"
  },
  {
    "id": 82,
    "productId": 123593,
    "supplierId": 16,
    "supplierMode": 2,
    "storeId": 30,
    "stock": 5,
    "deductionStock": 0,
    "isLastDeducted": 0,
    "date": "2024-05-29T00:00:00+08:00",
    "created": "2024-05-29T10:06:52+08:00",
    "updated": "2024-05-29T10:06:52+08:00"
  },
  {
    "id": 83,
    "productId": 123593,
    "supplierId": 24,
    "supplierMode": 1,
    "storeId": 30,
    "stock": 5,
    "deductionStock": 0,
    "isLastDeducted": 0,
    "date": "2024-05-29T00:00:00+08:00",
    "created": "2024-05-29T10:06:52+08:00",
    "updated": "2024-05-29T10:06:52+08:00"
  },
  {
    "id": 84,
    "productId": 123593,
    "supplierId": 25,
    "supplierMode": 1,
    "storeId": 30,
    "stock": 10,
    "deductionStock": 0,
    "isLastDeducted": 0,
    "date": "2024-05-29T00:00:00+08:00",
    "created": "2024-05-29T10:06:53+08:00",
    "updated": "2024-05-29T10:06:53+08:00"
  }
]
	`
	var logs []SupplierProductStockLog
	err := json.Unmarshal([]byte(jsonData), &logs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var purchaseSupplierSales, jointSupplierSales []*SupplierTotalSales
	for _, item := range logs {
		if item.SupplierMode == 1 {
			purchaseSupplierSales = aggregateSupplierSalesData(&item, purchaseSupplierSales)
		} else {
			jointSupplierSales = aggregateSupplierSalesData(&item, jointSupplierSales)
		}
	}
	prettyJson(purchaseSupplierSales)
}

func aggregateSupplierSalesData(item *SupplierProductStockLog,
	totalSales []*SupplierTotalSales,
) []*SupplierTotalSales {

	index := funk.IndexOf(totalSales, func(sale *SupplierTotalSales) bool {
		return item.SupplierId == sale.SupplierId
	})

	if index == -1 {
		totalSales = append(totalSales, &SupplierTotalSales{
			SalesQuantity: item.Stock,
			SupplierId:    item.SupplierId,
		})
	} else {
		targetSupplier := totalSales[index]
		targetSupplier.SalesQuantity += item.Stock
	}
	return totalSales
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
