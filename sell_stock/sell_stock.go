package sell_stock

import "math"

func MaxProfit(prices []int) int {
	// 初始化最低价格和最大利润
	minPrice, maxProfit := math.MaxInt64, 0

	for i := 0; i < len(prices); i++ {
		// 如果当天价格低于之前的最低价，设置该价格为最低买入价
		if prices[i] < minPrice {
			minPrice = prices[i]
			// 如果当天价格减去最低价 > 最大利润，最大利润就等于当天价格减去最低价
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}
