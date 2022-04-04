# 买卖股票的最佳时期

https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

## 题解

### 贪心算法

![](http://image.maplejoyous.cn/post/2022/04/04/202204041054033.png)

我们来假设自己来购买股票。随着时间的推移，每天我们都可以选择出售股票与否。那么，假设在第 i 天，如果我们要在今天卖股票，那么我们能赚多少钱呢？

显然，如果我们真的在买卖股票，我们肯定会想：如果我是在历史最低点买的股票就好了！太好了，在题目中，我们只要用一个变量记录一个历史最低价格 `minprice`，我们就可以假设自己的股票是在那天买的。那么我们在第 i 天卖出股票能得到的利润就是 `prices[i] - minprice`。

因此，我们只需要遍历价格数组一遍，记录历史最低点，然后在每一天考虑这么一个问题：如果我是在历史最低点买进的，那么我今天卖出能赚多少钱？当考虑完所有天数之时，我们就得到了最好的答案。

链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/121-mai-mai-gu-piao-de-zui-jia-shi-ji-by-leetcode-/

### DP 状态转移方程

考虑每次如何获取最大收益？第`i`天的最大收益只需要知道前`i`天的最低点就可以算出来了。而第`i`天以前（包括第 i 天）的最低点和`i-1`天的最低点有关，至此我们的动态方程就出来了。

```
dp[i] = min(dp[i-1], prices[i])
```

其中`dp[0]=prices[0]`,然后动态计算之后的就可以了。得到了前`i`天的最低点以后，只需要维护一个`maxProfit`用来保存最大收益就可以了。