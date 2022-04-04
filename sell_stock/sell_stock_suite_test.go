package sell_stock_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSellStock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SellStock Suite")
}
