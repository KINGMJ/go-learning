package simplyfactory

import "fmt"

// 简单工厂：SimplyFactory

const (
	ProductA = 0
	ProductB = 1
	ProductC = 2
)

func MakeProduct(product_kind int) IProduct {
	switch product_kind {
	case ProductA:
		return ConcreteProductA{}
	case ProductB:
		return ConcreteProductB{}
	default:
		return nil
	}
}

// 抽象产品
type IProduct interface {
	Show() // 展示商品
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------
// 具体产品 A
type ConcreteProductA struct{}

func (a ConcreteProductA) Show() {
	fmt.Println("I'm product A")
}

// 具体产品 A
type ConcreteProductB struct{}

func (b ConcreteProductB) Show() {
	fmt.Println("I'm product B")
}
