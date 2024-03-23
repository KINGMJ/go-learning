package factorymethod

import "fmt"

// 抽象工厂
type IAbstractFactory interface {
	MakeProduct() IProduct
}

// 抽象产品
type IProduct interface {
	Show() // 展示商品
}

// ----------- (●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●)(●'◡'●) ------------

// 具体工厂
type ProductAFactory struct{}

func (a ProductAFactory) MakeProduct() IProduct {
	return ConcreteProductA{}
}

type ProductBFactory struct{}

func (b ProductBFactory) MakeProduct() IProduct {
	return ConcreteProductB{}
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
