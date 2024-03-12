package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Code  string
	Price uint
}

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/vms2?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = database
}

func main() {
	createProduct()
}

// 自动迁移
func autoMigrate() {
	db.AutoMigrate(&Product{})
}

// 新增
func createProduct() {
	// 创建一条记录
	// result := db.Create(&Product{Name: "牙刷", Code: "PD-0002", Price: 4})
	// fmt.Println(result)

	// 通过指针创建
	// product := Product{Name: "杯子", Code: "PD-0003", Price: 6}
	// res := db.Create(&product)
	// fmt.Println(product)
	// fmt.Println(res.Error)
	// fmt.Println(res.RowsAffected)

	// 插入多条记录
	products := []Product{
		{Name: "香蕉", Code: "PD-0005", Price: 23},
		{Name: "桔子", Code: "PD-0006", Price: 32},
	}

	res := db.Create(&products)
	fmt.Println(products)
	fmt.Println(res.RowsAffected)

	// 批量插入
	// var products []Product
	// for i := 0; i < 200; i++ {
	// 	products = append(products, Product{
	// 		Name:  gofakeit.ProductName(),
	// 		Code:  "PD-000" + fmt.Sprintf("%d", i),
	// 		Price: uint(rand.Intn(100) + 1),
	// 	})
	// }
	// db.CreateInBatches(&products, 20)
	// for _, product := range products {
	// 	fmt.Println(product.ID)
	// }
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Code == "PD-0005" {
		return errors.New("invalid price ")
	}
	return
}
