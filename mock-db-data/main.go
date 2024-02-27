package main

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/sync/semaphore"
)

// 定义数据库对象
var Db *sqlx.DB

// 使用信号量限制每次并行执行5个任务
var (
	maxWorker = runtime.GOMAXPROCS(0)
	sem       = semaphore.NewWeighted(int64(maxWorker))
	weight    = 1
)

/*
初始化，连接db
*/
func init() {
	// 使用 dsn 的方式连接数据库
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mall")
	if err != nil {
		log.Fatal(err)
	}
	Db = database
	fmt.Println(Db) // &{0xc00010d380 mysql false 0xc00010a930}
}

func main() {
	insertTestData()
}

// 插入测试表
func insertTestData() {
	sqlDelete := "DELETE FROM t"
	Db.Exec(sqlDelete)
	sqlInsert := "INSERT INTO t(a,b) VALUES(?, ?)"
	for i := 1; i <= 100000; i++ {
		sem.Acquire(context.Background(), int64(weight))
		go func(i int) {
			defer sem.Release(int64(weight))
			result, err := Db.Exec(sqlInsert, i, i)
			if err != nil {
				log.Fatal(err)
			}
			id, err := result.LastInsertId()
			if err != nil {
				log.Fatal(fmt.Errorf("add t: %v", err))
			}
			fmt.Printf("ID of added t: %v\n", id)
		}(i)
	}
	if err := sem.Acquire(context.Background(), int64(maxWorker)); err != nil {
		log.Printf("获取所有的worker失败：%v", err)
	}
	fmt.Println("10w 条数据插入完毕...")
}

// 插入用户数据
func insertUserData() {
	// 插入 30w 条数据
	for i := 0; i < 300000; i++ {
		sem.Acquire(context.Background(), int64(weight))
		go func() {
			// mock 数据
			id, err := addUser(User{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				Phone: gofakeit.Phone(),
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("ID of added user: %v\n", id)
			sem.Release(int64(weight))
		}()
	}
	fmt.Println("30w 条数据插入完毕...")
}

type User struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Phone string `db:"phone"`
}

// insert
func addUser(user User) (int64, error) {
	//定义sql语句, 通过占位符 问号（ ? ) 定义了三个参数
	sqlInsert := "INSERT INTO user_test (name, email, phone) VALUES(?, ?, ?)"
	//通过Exec插入数据, 这里传入了三个参数，对应sql语句定义的三个问号所在的位置
	result, err := Db.Exec(sqlInsert, user.Name, user.Email, user.Phone)
	if err != nil {
		return 0, fmt.Errorf("add user: %v", err)
	}
	//插入成功后，获取insert id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add user: %v", err)
	}
	return id, nil
}
