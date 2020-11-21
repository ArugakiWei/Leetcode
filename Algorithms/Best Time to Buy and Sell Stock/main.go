package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// dp[i][1] 第i天, 持有股票
	// dp[i][0] 第i天, 未持有股票
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}

	// 第0天未持有股票, 利润是0
	dp[0][0] = 0
	// 第0天持有股票, 当天购买
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		// 第i天没有股票=(i-1天就没有股票, i-1天有股票, i天卖了)
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 第i天有股票=(i天就有股票, i-1天没有股票, i天买了)
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	fmt.Println(dp)
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
