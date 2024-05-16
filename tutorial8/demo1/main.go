package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx/v3"
)

func main() {
	xlsxRead()
}

type Person struct {
	Name  string
	Age   int
	Email string
}

func xlsxRead() {
	wb, err := xlsx.OpenFile("./data.xlsx")
	if err != nil {
		log.Fatalf("无法打开文件：%v", err)
	}
	// 获取sheet
	sheet := wb.Sheets[0]

	// max Row
	fmt.Println("Max row is", sheet.MaxRow)
	// 遍历每一行（跳过标题）
	err = sheet.ForEachRow(rowVisitor)
	fmt.Println("Err=", err)
}

func rowVisitor(r *xlsx.Row) error {
	// 每一行中迭代单元格
	return r.ForEachCell(cellVisitor)
}

func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Cell value:", value)
	}
	return err
}
