package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mitchellh/mapstructure"
)

// 定义数据库对象
var Db *sqlx.DB

type Album struct {
	ID     int64   `db:"id"`
	Title  string  `db:"title"`
	Artist string  `db:"artist"`
	Price  float32 `db:"price"`
}

/*
初始化，连接db
*/
func init() {
	// 使用 dsn 的方式连接数据库
	database, err := sqlx.Open("mysql", "root:leangoo123@tcp(127.0.0.1:3306)/recording")
	if err != nil {
		log.Fatal(err)
	}
	Db = database
	fmt.Println(Db) // &{0xc00010d380 mysql false 0xc00010a930}
}

func main() {
	selectDemo()
}

func selectDemo() {
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(albums)
	fmt.Print(albums[0])

	album, err := albumsById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(album)
}

func insertDemo() {
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
}

func updateDemo() {
	row, err := updateAlbum(3, "Jeruex")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(row)
}

func deleteDemo() {
	row, err := deleteAlbum(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(row)
}

func queryXDemo() {
	albums, err := getAllAlbums()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(albums)
}

// select 操作，获取多行
func albumsByArtist(name string) ([]Album, error) {
	var album []Album
	err := Db.Select(&album, "SELECT * FROM album WHERE artist= ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return album, nil
}

// get 操作，获取单行
func albumsById(id int64) (Album, error) {
	// 定义保存查询结果的变量
	var album Album
	// 通过指针将查询结果保存在 album 变量中
	if err := Db.Get(&album, "SELECT * FROM album WHERE id= ?", id); err != nil {
		return album, fmt.Errorf("albumsById %d: no such album", id)
	}
	return album, nil
}

// query 操作，通过不同方式获取数据
func getAllAlbums() ([]Album, error) {
	// 查询所有的数据，这里返回的是sqlx.Rows对象
	rows, err := Db.Queryx("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("getAllAlbums: no data")
	}
	// albums, err := scanRows(rows)
	// albums, err := structScanRows(rows)
	// albums, err := mapScanRows(rows)
	albums, err := sliceScanRows(rows)
	if err != nil {
		return nil, fmt.Errorf("getAllAlbums: no data")
	}
	return albums, nil
}

// 通过 Scan 获取数据
func scanRows(rows *sqlx.Rows) ([]Album, error) {
	// 定义一个 slice，用于接收数据
	var albums []Album
	//这里定义4个变量用于接收每一行数据
	var id int64
	var title string
	var artist string
	var price float32
	// 循环遍历每一行记录，rows.Next()函数用于判断是否还有下一行数据
	for rows.Next() {
		// 通过 Scan 函数将结果保存在变量中
		err := rows.Scan(&id, &title, &artist, &price)
		if err != nil {
			return nil, err
		}
		album := Album{id, title, artist, price}
		albums = append(albums, album)
	}
	return albums, nil
}

// 通过 StructScan 获取数据
func structScanRows(rows *sqlx.Rows) ([]Album, error) {
	var albums []Album
	// 定义一个 struct，用于接收数据
	var album Album
	for rows.Next() {
		err := rows.StructScan(&album)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

// 通过 MapScan 获取数据
func mapScanRows(rows *sqlx.Rows) ([]Album, error) {
	var albums []Album
	var album Album
	//定义 map 类型
	m := make(map[string]interface{})
	for rows.Next() {
		err := rows.MapScan(m)
		if err != nil {
			return nil, err
		}
		// MapScan 返回的结果是二进制的，需要转换为 string
		for k, v := range m {
			if _, ok := v.([]byte); ok {
				m[k] = string(v.([]byte))
			}
		}
		// 将 map 转换为 struct
		if err := mapstructure.WeakDecode(m, &album); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

// 通过 SliceScan 获取数据
func sliceScanRows(rows *sqlx.Rows) ([]Album, error) {
	for rows.Next() {
		s, err := rows.SliceScan()
		if err != nil {
			return nil, err
		}
		for idx, v := range s {
			if _, ok := v.([]byte); ok {
				s[idx] = string(v.([]byte))
			}
		}
		fmt.Println(s)
		/*
			[1 Blue Train John Coltrane 56.99]
			[2 Giant Steps John Coltrane 63.99]
			[4 Sarah Vaughan Sarah Vaughan 34.98]
			[5 The Modern Sound of Betty Carter Betty Carter 49.99]
			[6 The Modern Sound of Betty Carter Betty Carter 49.99]
			[]
		*/
	}
	return nil, nil
}

// insert
func addAlbum(alb Album) (int64, error) {
	//定义sql语句, 通过占位符 问号（ ? ) 定义了三个参数
	sqlInsert := "INSERT INTO album (title, artist, price) VALUES(?, ?, ?)"
	//通过Exec插入数据, 这里传入了三个参数，对应sql语句定义的三个问号所在的位置
	result, err := Db.Exec(sqlInsert, alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	//插入成功后，获取insert id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

// 更新操作
func updateAlbum(id int64, title string) (int64, error) {
	result, err := Db.Exec("UPDATE album set title= ? WHERE id= ?", title, id)
	if err != nil {
		return 0, fmt.Errorf("updateAlbum: %v", err)
	}
	// 查询更新影响的行数
	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateAlbum: %v", err)
	}
	return row, nil
}

// 删除操作
func deleteAlbum(id int64) (int64, error) {
	result, err := Db.Exec("DELETE FROM album WHERE id= ?", id)
	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %v", err)
	}
	row, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %v", err)
	}
	return row, nil
}
