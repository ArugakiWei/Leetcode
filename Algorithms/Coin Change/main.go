package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange([]int{1}, 1))
	fmt.Println(coinChange([]int{1}, 2))
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	dp := make([]int, amount + 1)
	for i:=1; i<len(dp); i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if i - coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin] + 1)
		}
	}

	if dp[amount] != math.MaxInt32 {
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}