package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	K := 2
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < len(prices); i++ {
		for k := 1; k <= K; k++ {
			if i == 0 {
				dp[0][k][0] = 0
				dp[0][k][1] = -prices[0]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[len(prices)-1][K][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
