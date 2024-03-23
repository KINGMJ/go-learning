package simplyfactory_test

import (
	"fmt"
	simplyfactory "oop/simply_factory"
	"testing"
)

func TestSimplyFactory(t *testing.T) {

	fmt.Println("开始创建产品...")

	productA := simplyfactory.MakeProduct(simplyfactory.ProductA)
	productA.Show()

	productB := simplyfactory.MakeProduct(simplyfactory.ProductB)
	productB.Show()

}
