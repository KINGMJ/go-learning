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
