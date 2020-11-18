package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		if i > 0 {
			dp[i][0] = i
		}
	}
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}
