package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/KINGMJ/go-learning/tutorial11/amount"
	restyclient "github.com/KINGMJ/go-learning/tutorial13/resty_client"
)

func main() {
	demo2()
}

type PurchaseOrderListReq struct {
	StartDate           string `json:"startDate" validate:"c_date"`                  // 开始日期
	EndDate             string `json:"endDate" validate:"c_date"`                    // 结束日期
	Page                int64  `json:"page" validate:"min=1"`                        // 页码，最小值为1
	Size                int64  `json:"size" validate:"min=1"`                        // 每页大小，最小值为1
	Status              int64  `json:"orderStatus,optional,options=[1,2,3,4,5,6,7]"` // 采购单状态：1 待提交；2 待接单；3 已取消；4 已拒绝；5 待发货；6 已发货；7 已收货
	RecommendSupplierID int64  `json:"supplierId,optional"`                          // 推荐供应商ID
}

func demo1() {
	purchaseListReq := PurchaseOrderListReq{
		StartDate: "2024-01-01",
		EndDate:   "2024-07-15",
		Page:      1,
		Size:      10,
		Status:    1,
	}
	var res restyclient.Response[Pagination[PurchaseOrderWithExtra]]
	err := restyclient.Post("http://localhost:9292/purchase-order/list", purchaseListReq, &res)
	if err != nil {
		log.Fatalf("%v", err)
	}
	prettyJson(res.Data)
}

func demo2() {
	// 创建采购单
	reqStr := `
	{
    "storeId": 30,
    "recommendSupplierId": 1,
    "shippingAddress": "这是一个门店地址",
    "expectedTime": "2024-01-02 12:12:11",
    "payType": 1,
    "products": [
        {
            "id": 2,
            "quantity": 100
        },
        {
            "id": 3,
            "quantity": 100
        }
    ],
    "boothSpace": "30",
    "warehouseSpace": "10",
    "stackSpace": "20"
}
	`
	var req PurchaseOrderCreateReq
	err := json.Unmarshal([]byte(reqStr), &req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var res restyclient.Response[PurchaseOrderCreateRes]
	err = restyclient.Post("http://localhost:9292/purchase-order/create", req, &res)
	if err != nil {
		log.Fatalf("%v", err)
	}
	prettyJson(res.Data)

}

type PurchaseOrderCreateReq struct {
	StoreID             int64                  `json:"storeId"`                            // 门店ID
	RecommendSupplierID int64                  `json:"recommendSupplierId"`                // 推荐供应商ID
	ExpectedTime        string                 `json:"expectedTime" validate:"c_datetime"` // 期望送达时间
	PayType             int                    `json:"payType"`                            // 付款方式：1 先货后款；2 线下付款
	ShippingAddress     string                 `json:"shippingAddress"`                    // 收货地址
	Products            []PurchaseOrderProduct `json:"products"`                           // 采购单商品
	BoothSpace          amount.Amount          `json:"boothSpace"`                         // 展位空间
	WarehouseSpace      amount.Amount          `json:"warehouseSpace"`                     // 仓储空间
	StackSpace          amount.Amount          `json:"stackSpace"`                         // 堆头空间
}

type PurchaseOrderCreateRes struct {
	PurchaseOrderID int64 `json:"purchaseOrderId"` // 采购单ID
}

type PurchaseOrderProduct struct {
	ID       int64 `json:"id"`       // 商品Id
	Quantity int64 `json:"quantity"` // 采购件数
}

type Pagination[T any] struct {
	Items []*T  `json:"items"` // 数据
	Total int64 `json:"total"` // 总页数
	Page  int64 `json:"page"`  // 当前页
	Size  int64 `json:"size"`  // 数据条数
}

func New[T any](page, size int64) *Pagination[T] {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	return &Pagination[T]{
		Page:  page,
		Size:  size,
		Items: make([]*T, 0), // 在这个地方声明，避免查询不到数据时，返回null的问题
	}
}

// Offset 获取数据库查询的偏移量
func (p *Pagination[T]) Offset() int64 {
	return (p.Page - 1) * p.Size
}

func (p *Pagination[T]) Limit() int64 {
	return p.Size
}

type PurchaseOrderWithExtra struct {
	PurchaseOrder
	SupplierName     string                      `db:"supplier_name" json:"supplierName"`          // 推荐供应商名称
	SupplierTypeName string                      `db:"supplier_type_name" json:"supplierTypeName"` // 推荐供应商分类名称
	StoreSpace       PurchaseOrderSpaceWithExtra `json:"storeSpace"`                               // 采购单空间
}

type PurchaseOrder struct {
	ID                  int64  `db:"id" json:"id" goqu:"skipinsert,skipupdate"`        // Id
	MerchantID          int64  `db:"merchant_id" json:"merchantId"`                    // 商户Id
	StoreID             int64  `db:"store_id" json:"storeId"`                          // 门店Id
	RecommendSupplierID int64  `db:"recommend_supplier_id" json:"recommendSupplierId"` // 推荐供应商Id
	No                  string `db:"no" json:"no" goqu:"skipupdate"`                   // 采购单号
	SalesNo             string `db:"sales_no" json:"salesNo"`                          // 销售单号
	TracksNo            string `db:"tracks_no" json:"tracksNo"`                        // 物流单号，多个单号用逗号分隔
	// ExpectedTime        time.Time     `db:"expected_time" json:"expectedTime" goqu:"omitempty"`  // 期望送达时间
	// SubmitTime          time.Time     `db:"submit_time" json:"submitTime" goqu:"omitempty"`      // 提交时间
	PayType         int           `db:"pay_type" json:"payType"`                 // 付款方式：1 先货后款；2 线下付款
	ShippingAddress string        `db:"shipping_address" json:"shippingAddress"` // 收货地址
	Count           int64         `db:"count" json:"purchaseCount"`              // 商品总数量
	Amount          amount.Amount `db:"amount" json:"purchaseAmount"`            // 商品总金额
	Status          int           `db:"status" json:"status"`                    // 采购单状态：1 待提交；2 待接单；3 已取消；4 已拒绝；5 待发货；6 已发货；7 已收货
	// Created         time.Time     `db:"created" json:"created" goqu:"skipinsert,skipupdate"` // 创建时间
	// Updated         time.Time     `db:"updated" json:"updated" goqu:"skipinsert,skipupdate"` // 更新时间
}

type PurchaseOrderSpaceWithExtra struct {
	PurchaseOrderSpace
	MerchantName string `db:"merchant_name" json:"merchantName"` // 商户名称
	StoreName    string `db:"store_name" json:"storeName"`       // 门店名称
}

type PurchaseOrderSpace struct {
	ID              int64         `db:"id" json:"id" goqu:"skipinsert,skipupdate"` // Id
	PurchaseOrderID int64         `db:"purchase_order_id" json:"purchaseOrderId"`  // 采购单Id
	BoothSpace      amount.Amount `db:"booth_space" json:"boothSpace"`             // 展位空间
	BoothAmount     amount.Amount `db:"booth_amount" json:"boothAmount"`           // 展位收费
	WarehouseSpace  amount.Amount `db:"warehouse_space" json:"warehouseSpace"`     // 仓储空间
	WarehouseAmount amount.Amount `db:"warehouse_amount" json:"warehouseAmount"`   // 仓储收费
	StackSpace      amount.Amount `db:"stack_space" json:"stackSpace"`             // 堆头空间
	StackAmount     amount.Amount `db:"stack_amount" json:"stackAmount"`           // 堆头收费
	// Created         time.Time     `db:"created" json:"created" goqu:"skipinsert,skipupdate"` // 创建时间
	// Updated         time.Time     `db:"updated" json:"updated" goqu:"skipinsert,skipupdate"` // 更新时间
}

func prettyJson(req any) {
	data, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling to JSON: %v", err)
	}
	prettyOutput := (string(data))
	fmt.Println(prettyOutput)
}
