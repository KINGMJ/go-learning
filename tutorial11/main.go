package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/KINGMJ/go-learning/tutorial11/amount"
	mTime "github.com/KINGMJ/go-learning/tutorial11/time"
)

type PurchaseSaleListRes struct {
	SupplierPurchaseSaleLog
	SupplierName string `db:"supplier_name" json:"supplierName"`
	StoreName    string `db:"store_name" json:"storeName"`
}

type SupplierPurchaseSaleLog struct {
	Id              int64         `db:"id" json:"id" goqu:"skipinsert,skipupdate"`                     // 编号ID
	MerchantId      int64         `db:"merchant_id" json:"merchantId"`                                 // 商户ID
	SupplierId      int64         `db:"supplier_id" json:"supplierId"`                                 // 供应商ID
	StoreId         int64         `db:"store_id" json:"storeId"`                                       // 门店ID
	InQuantity      int           `db:"in_quantity" json:"inQuantity"`                                 // 入库数量
	InAmount        amount.Amount `db:"in_amount" json:"inAmount"`                                     // 入库贷款
	SettledQuantity int           `db:"settled_quantity" json:"settledQuantity"`                       // 已结数量
	SettledAmount   amount.Amount `db:"settled_amount" json:"settledAmount"`                           // 已结贷款
	SettledFee      amount.Amount `db:"settled_fee" json:"settledFee"`                                 // 已结杂项
	SettledPayable  amount.Amount `db:"settled_payable" json:"settledPayable"`                         // 已结应付
	SalesQuantity   int           `db:"sales_quantity" json:"salesQuantity"`                           // 销售数量
	SalesAmount     amount.Amount `db:"sales_amount" json:"salesAmount"`                               // 销售金额
	Date            time.Time     `db:"date" json:"date" goqu:"omitempty,skipupdate"`                  // 日期
	Created         time.Time     `db:"created" json:"created" goqu:"omitempty,skipinsert,skipupdate"` // 创建时间
	Updated         time.Time     `db:"updated" json:"updated" goqu:"omitempty,skipinsert,skipupdate"` // 修改时间
}

func demo1() {
	jsonData := `
	[
            {
                "id": 76,
                "merchantId": 15,
                "supplierId": 14,
                "storeId": 30,
                "inQuantity": 58,
                "inAmount": "430.00",
                "settledQuantity": 18,
                "settledAmount": "90.00",
                "settledFee": "0.00",
                "settledPayable": "90.00",
                "salesQuantity": 38,
                "salesAmount": "260.00",
                "date": "2024-01-19 00:00:00",
                "created": "2024-05-30 11:40:09",
                "updated": "2024-05-30 11:40:09",
                "supplierName": "东方供应链有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 73,
                "merchantId": 15,
                "supplierId": 14,
                "storeId": 30,
                "inQuantity": 55,
                "inAmount": "415.00",
                "settledQuantity": 15,
                "settledAmount": "75.00",
                "settledFee": "0.00",
                "settledPayable": "75.00",
                "salesQuantity": 35,
                "salesAmount": "245.00",
                "date": "2024-01-16 00:00:00",
                "created": "2024-05-30 11:40:09",
                "updated": "2024-05-30 11:40:09",
                "supplierName": "东方供应链有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 64,
                "merchantId": 15,
                "supplierId": 14,
                "storeId": 30,
                "inQuantity": 46,
                "inAmount": "370.00",
                "settledQuantity": 6,
                "settledAmount": "30.00",
                "settledFee": "0.00",
                "settledPayable": "30.00",
                "salesQuantity": 26,
                "salesAmount": "200.00",
                "date": "2024-01-07 00:00:00",
                "created": "2024-05-30 11:40:09",
                "updated": "2024-05-30 11:40:09",
                "supplierName": "东方供应链有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 70,
                "merchantId": 15,
                "supplierId": 14,
                "storeId": 30,
                "inQuantity": 52,
                "inAmount": "400.00",
                "settledQuantity": 12,
                "settledAmount": "60.00",
                "settledFee": "0.00",
                "settledPayable": "60.00",
                "salesQuantity": 32,
                "salesAmount": "230.00",
                "date": "2024-01-13 00:00:00",
                "created": "2024-05-30 11:40:09",
                "updated": "2024-05-30 11:40:09",
                "supplierName": "东方供应链有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 67,
                "merchantId": 15,
                "supplierId": 14,
                "storeId": 30,
                "inQuantity": 49,
                "inAmount": "385.00",
                "settledQuantity": 9,
                "settledAmount": "45.00",
                "settledFee": "0.00",
                "settledPayable": "45.00",
                "salesQuantity": 29,
                "salesAmount": "215.00",
                "date": "2024-01-10 00:00:00",
                "created": "2024-05-30 11:40:09",
                "updated": "2024-05-30 11:40:09",
                "supplierName": "东方供应链有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 1,
                "merchantId": 15,
                "supplierId": 14,
                "storeId": 30,
                "inQuantity": 40,
                "inAmount": "340.00",
                "settledQuantity": 0,
                "settledAmount": "0.00",
                "settledFee": "0.00",
                "settledPayable": "0.00",
                "salesQuantity": 20,
                "salesAmount": "170.00",
                "date": "2024-05-29 00:00:00",
                "created": "2024-05-29 17:15:02",
                "updated": "2024-05-29 18:06:34",
                "supplierName": "东方供应链有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 68,
                "merchantId": 15,
                "supplierId": 24,
                "storeId": 30,
                "inQuantity": 50,
                "inAmount": "390.00",
                "settledQuantity": 10,
                "settledAmount": "50.00",
                "settledFee": "0.00",
                "settledPayable": "50.00",
                "salesQuantity": 30,
                "salesAmount": "220.00",
                "date": "2024-01-11 00:00:00",
                "created": "2024-05-30 11:40:09",
                "updated": "2024-05-30 11:40:09",
                "supplierName": "西部综合贸易有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 77,
                "merchantId": 15,
                "supplierId": 24,
                "storeId": 30,
                "inQuantity": 59,
                "inAmount": "435.00",
                "settledQuantity": 19,
                "settledAmount": "95.00",
                "settledFee": "0.00",
                "settledPayable": "95.00",
                "salesQuantity": 39,
                "salesAmount": "265.00",
                "date": "2024-01-20 00:00:00",
                "created": "2024-05-30 11:40:09",
                "updated": "2024-05-30 11:40:09",
                "supplierName": "西部综合贸易有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 74,
                "merchantId": 15,
                "supplierId": 24,
                "storeId": 30,
                "inQuantity": 56,
                "inAmount": "420.00",
                "settledQuantity": 16,
                "settledAmount": "80.00",
                "settledFee": "0.00",
                "settledPayable": "80.00",
                "salesQuantity": 36,
                "salesAmount": "250.00",
                "date": "2024-01-17 00:00:00",
                "created": "2024-05-30 11:40:09",
                "updated": "2024-05-30 11:40:09",
                "supplierName": "西部综合贸易有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 71,
                "merchantId": 15,
                "supplierId": 24,
                "storeId": 30,
                "inQuantity": 53,
                "inAmount": "405.00",
                "settledQuantity": 13,
                "settledAmount": "65.00",
                "settledFee": "0.00",
                "settledPayable": "65.00",
                "salesQuantity": 33,
                "salesAmount": "235.00",
                "date": "2024-01-14 00:00:00",
                "created": "2024-05-30 11:40:09",
                "updated": "2024-05-30 11:40:09",
                "supplierName": "西部综合贸易有限公司",
                "storeName": "莱茵零售店"
            },
            {
                "id": 0,
                "merchantId": 0,
                "supplierId": 0,
                "storeId": 0,
                "inQuantity": 904,
                "inAmount": "7040.00",
                "settledQuantity": 184,
                "settledAmount": "920.00",
                "settledFee": "0.00",
                "settledPayable": "920.00",
                "salesQuantity": 564,
                "salesAmount": "4150.00",
                "date": "0001-01-01 00:00:00",
                "created": "0001-01-01 00:00:00",
                "updated": "0001-01-01 00:00:00",
                "supplierName": "",
                "storeName": ""
            }
        ]
	`
	var logs []*PurchaseSaleListRes
	err := json.Unmarshal([]byte(jsonData), &logs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 最后结果集合
	var resList []*PurchaseSaleListRes
	// 每个供应商的聚合数据
	summaryData := &PurchaseSaleListRes{}
	var prevSupplierId int64 = 0

	for _, item := range logs[:len(logs)-1] {
		if item.SupplierId != prevSupplierId && prevSupplierId > 0 {
			resList = append(resList, summaryData)
			// 重置summary
			summaryData = &PurchaseSaleListRes{}
		}

		summaryData.SupplierName = item.SupplierName
		summaryData.MerchantId = item.MerchantId
		summaryData.SupplierId = item.SupplierId

		// 累加数据
		summaryData.InQuantity += item.InQuantity
		summaryData.InAmount += item.InAmount
		summaryData.SettledQuantity += item.SettledQuantity
		summaryData.SettledAmount += item.SettledAmount
		summaryData.SettledFee += item.SettledFee
		summaryData.SettledPayable += item.SettledPayable
		summaryData.SalesQuantity += item.SalesQuantity
		summaryData.SalesAmount += item.SalesAmount

		prevSupplierId = item.SupplierId
		resList = append(resList, item)
	}
	resList = append(resList, summaryData)
	resList = append(resList, logs[len(logs)-1])
	prettyJson(resList)
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

func main() {
	t := time.Date(2024, 6, 1, 12, 31, 0, 0, time.UTC)

	t1 := mTime.Time{t}
	t2 := mTime.Time{time.Now()}

	fmt.Println(ComputeDiffDays(t2.Time, t1.Time))

}

func ComputeDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours()) / 24
}
