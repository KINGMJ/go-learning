package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

// 数据库的配置
var (
	dsn string = "root:123456@tcp(localhost:3306)/leangoo?charset=utf8mb4&parseTime=True&loc=Local"
	db  *gorm.DB
)

type Task struct {
	gorm.Model
	TaskID              string  `gorm:"column:task_id;type:char(16);not null"`
	TaskName            string  `gorm:"column:task_name;type:varchar(4000);not null"`
	BoardID             uint    `gorm:"column:board_id;not null"`
	ListID              uint    `gorm:"column:list_id"`
	Status              int32   `gorm:"column:status;not null;default:1"`
	Estimate            float64 `gorm:"column:estimate"`
	Type                uint    `gorm:"column:type;not null;default:1"`
	EstimateWorkingHour float64 `gorm:"column:estimate_working_hour;type:float(10,2)"`
}

// 设置表名
func (Task) TableName() string {
	return "lg_task"
}

func init() {
	// 使用默认连接方式
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = database
}

func main() {
	// 数据迁移，将现有的lg_task表进行水平拆分
	// shardTaskMigrate()

	// hash 取模拆分
	hashSharedMigrate()
}

// 范围分配，根据 board_id 作为水平拆分的健进行范围拆分
func rangeSharedMigrate() {
	// var task Task
	var rowCount int = 400000 // 每个表的容量
	// db.Migrator().DropTable(&Task{})
	// db.Table(task.TableName()).AutoMigrate(&Task{})

	// 1. 我们查出当前表的总数量大概为 400w 条记录，我们可以按照每个表40w条，分成11个表（因为总条数>400w）
	for i := 0; i < 11; i++ {
		tableName := fmt.Sprintf("lg_task_%0*d", 3, i)
		db.Table(tableName).AutoMigrate(&Task{})

		// 2. 进行数据迁移
		sql := `INSERT INTO %s (task_id, task_name, board_id, list_id,
			create_datetime, status, estimate, type, update_datetime, estimate_working_hour)
	 SELECT task_id, task_name, board_id, list_id, create_datetime, status, estimate, type,
	  update_datetime, estimate_working_hour
	 FROM lg_task
	 ORDER BY board_id ASC
	 LIMIT %d, %d`
		// 按照每个表40w进行迁移
		sql = fmt.Sprintf(sql, tableName, rowCount*i, rowCount)
		res := db.Exec(sql)
		if res.Error != nil {
			log.Fatal(res.Error)
		}
	}
}

// 使用 gorm PKSnowflake 算法进行哈希水平拆分
func hashSharedMigrate() {
	// 分表，创建64个表
	for i := 0; i < 12; i++ {
		tableName := fmt.Sprintf("lg_task_%0*d", 2, i)
		db.Table(tableName).AutoMigrate(&Task{})
	}

	// 配置PKSnowflake拆分
	db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "task_id",
		NumberOfShards:      12, // 分为64个表
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "lg_task"))

	db.Create(&Task{
		TaskID:   "0fbc28f9ba733714",
		TaskName: "xx",
		BoardID:  1,
		ListID:   919177,
		Status:   1,
		Type:     1,
	})
}
