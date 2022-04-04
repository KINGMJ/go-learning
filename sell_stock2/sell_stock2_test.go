package sell_stock2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"go-learning/sell_stock2"
)

var MaxProfit = sell_stock2.MaxProfit2

var _ = Describe("SellStock2", func() {
	It("test1", func() {
		Expect(MaxProfit([]int{7, 1, 5, 3, 6, 4})).To(Equal(7))
	})

	It("test2", func() {
		Expect(MaxProfit([]int{1, 2, 3, 4, 5})).To(Equal(4))
	})

	It("test3", func() {
		Expect(MaxProfit([]int{7, 6, 4, 3, 1})).To(Equal(0))
	})
})
