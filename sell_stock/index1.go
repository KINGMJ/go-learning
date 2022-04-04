package sell_stock

func MaxProfit1(prices []int) int {
	var maxProfit int = 0
	// 使用dp 状态转移方程
	dp := make([]int, len(prices))
	dp[0] = prices[0]

	for i := 1; i < len(prices); i++ {
		if dp[i-1] < prices[i] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = prices[i]
		}

		if prices[i]-dp[i] > maxProfit {
			maxProfit = prices[i] - dp[i]
		}
	}
	return maxProfit
}

/*
 状态转移方程： dp[i] = min(dp[i-1], price[i])
*/
