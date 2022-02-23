package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

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
	database, err := sqlx.Open("mysql", "root:leangoo123@tcp(127.0.0.1:3306)/recording")
	if err != nil {
		log.Fatal(err)
	}
	Db = database
	fmt.Println(Db) // &{0xc00010d380 mysql false 0xc00010a930}
	// defer database.Close()
}

func main() {
	deleteDemo()
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

	getAllAlbums()
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
	var album Album
	if err := Db.Get(&album, "SELECT * FROM album WHERE id= ?", id); err != nil {
		return album, fmt.Errorf("albumsById %d: no such album", id)
	}
	return album, nil
}

// query 操作
func getAllAlbums() {
	var album Album
	rows, err := Db.Queryx("SELECT * FROM album")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.StructScan(&album)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", album)
	}
}

// insert
func addAlbum(alb Album) (int64, error) {
	result, err := Db.Exec("INSERT INTO album (title, artist, price) VALUES(?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
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
