package main

import "fmt"

func main() {
	fmt.Println(longestIdealString("acfgbd", 2))
	fmt.Println(longestIdealString("acbd", 2))
	fmt.Println(longestIdealString("eduktdb", 15))
}

func longestIdealString(s string, k int) int {
	dp := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cur := int(s[i] - 'a')
		curMax := 0
		minReachLetter := max(cur-k, 0)
		maxReachLetter := min(cur+k, 25)
		for c := minReachLetter; c <= maxReachLetter; c++ {
			if dp[c] > curMax {
				curMax = dp[c]
			}
		}
		dp[cur] = curMax + 1
	}

	ans := 0
	for _, v := range dp {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
