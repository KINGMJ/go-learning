package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// 数据库的配置
var (
	dbSeparation bool     = true // 是否使用读写分离
	dsnMaster    string   = "root:123456@tcp(192.168.3.36:3306)/vms2?charset=utf8mb4&parseTime=True&loc=Local"
	dsnSlave     []string = []string{
		"root:123456@tcp(192.168.3.31:3306)/vms2?charset=utf8mb4&parseTime=True&loc=Local",
		"root:123456@tcp(192.168.3.48:3306)/vms2?charset=utf8mb4&parseTime=True&loc=Local",
		"root:123456@tcp(192.168.3.53:3306)/vms2?charset=utf8mb4&parseTime=True&loc=Local",
	}
)

type Product struct {
	gorm.Model
	Name  string
	Code  string
	Price uint
}

var db *gorm.DB

func init() {
	// 使用读写分离连接方式
	if dbSeparation {
		connectRWDB()
		return
	}
	// 使用默认连接方式
	database, err := gorm.Open(mysql.Open(dsnMaster), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = database
}

func connectRWDB() {
	database, err := gorm.Open(mysql.Open(dsnMaster), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// 配置从库
	replicas := []gorm.Dialector{}
	for _, dsn := range dsnSlave {
		replicas = append(replicas, mysql.Open(dsn))
	}
	database.Use(
		dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dsnMaster)},
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		}),
	)
	db = database
}

func main() {
	selectProduct()
}

func selectProduct() {
	var product Product
	// db.First(&product)

	// 手动切换连接
	db.Clauses(dbresolver.Write).First(&product)

	data, _ := json.MarshalIndent(product, "", " ")
	fmt.Printf("%s\n", data)

}
