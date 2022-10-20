package main

import "fmt"

func main() {
	fmt.Println(longestArithSeqLength([]int{1, 2, 3, 4}))
}

func longestArithSeqLength(nums []int) int {
	dp := make([]map[int]int, len(nums))
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	diff, max := 0, 0
	for i, num := range nums {
		for j := 0; j < i; j++ {
			diff = num - nums[j]
			if cnt, ok := dp[j][diff]; ok {
				dp[i][diff] = cnt + 1
			} else {
				dp[i][diff] = 2
			}
			if max < dp[i][diff] {
				max = dp[i][diff]
			}
		}
	}
	return max
}
