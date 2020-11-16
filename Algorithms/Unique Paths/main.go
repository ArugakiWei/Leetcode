package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, n+1)
		if i == 1 {
			for j := 0; j <= n; j++ {
				dp[i][j] = 1
			}
		}
		dp[i][1] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}
