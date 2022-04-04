package sell_stock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go-learning/sell_stock"
)

var MaxProfit = sell_stock.MaxProfit1

var _ = Describe("SellStock", func() {
	It("test1", func() {
		Expect(MaxProfit([]int{7, 1, 5, 3, 6, 4})).To(Equal(5))
	})

	It("test2", func() {
		Expect(MaxProfit([]int{7, 6, 4, 3, 1})).To(Equal(0))
	})

	It("test3", func() {
		Expect(MaxProfit([]int{2, 4, 1})).To(Equal(2))
	})
})
