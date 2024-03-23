package factorymethod_test

import (
	factorymethod "oop/factory_method"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	factoryA := factorymethod.ProductAFactory{}
	factoryA.MakeProduct().Show()

	factoryB := factorymethod.ProductBFactory{}
	factoryB.MakeProduct().Show()
}
