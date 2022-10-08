package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump1(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 1; i < len(dp); i++ {
		for j := 1; j <= i; j++ {
			if nums[i-j] >= j {
				if dp[i] == 0 {
					dp[i] = dp[i-j] + 1
				} else {
					dp[i] = min(dp[i], dp[i-j]+1)
				}
			}
		}
	}
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func jump(nums []int) int {
	maxIndex, currentMaxIndex, result := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxIndex = max(maxIndex, i+nums[i])
		if i == currentMaxIndex {
			currentMaxIndex = maxIndex
			result++
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
