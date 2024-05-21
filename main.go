package main

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

func main() {
	sql, _, _ := goqu.From("test").ToSQL()

	fmt.Println(sql)

}
