package main

import "fmt"

func main() {
	// fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))           // 84
	// fmt.Println(maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4)) // 83
	// fmt.Println(maxSumAfterPartitioning([]int{8, 5, 7}, 3))                         // 24
	fmt.Println(maxSumAfterPartitioning([]int{10, 9, 3, 2}, 2)) // 30
}

func maxSumAfterPartitioning(arr []int, k int) int {
	maxNum := func(i, j int) int {
		maxx := arr[i]
		for index := i; index <= j; index++ {
			if arr[index] > maxx {
				maxx = arr[index]
			}
		}
		return maxx
	}
	dp := make([]int, len(arr))
	for i := 0; i < k; i++ {
		dp[i] = maxNum(0, i) * (i + 1)
	}
	for i := k; i < len(arr); i++ {
		for j := 1; j <= k; j++ {
			dp[i] = max(dp[i], dp[i-j]+maxNum(i-j+1, i)*j)
		}
	}
	return dp[len(arr)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
