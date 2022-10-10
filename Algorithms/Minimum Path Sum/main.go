package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	// dp[i][j] 表示到 点(i,j) 的最短路径和
	// 初始化当 grid 只有1行和1列的情况
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
		if i == 0 {
			dp[i][0] = grid[i][0]
			for j := 1; j < len(grid[i]); j++ {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}

	// 到点(i,j)只能是点(i-1,j)向下或者点(i,j-1)向右移动一步
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
