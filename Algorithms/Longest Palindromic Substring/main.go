package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("bbbbb"))
	fmt.Println(longestPalindrome("abcda"))
	fmt.Println(longestPalindrome("aacabdkacaa"))
}

func longestPalindrome1(s string) string {
	// s[i..j] 中最长回文子串的长度
	dp := make([][]int, len(s))
	for i, _ := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	dp[0][0] = 0

	ss := []byte(s)
	res, start, end := 0, 0, 0
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			// 如果头尾相等
			if ss[i] == ss[j] {
				// 如果去掉头尾仍是回文串
				if dp[i+1][j-1] == j-i-1 {
					dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)

					if dp[i][j] > res {
						res = dp[i][j]
						start = i
						end = j
					}
				}
			}
		}
	}

	return string(ss[start : end+1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var res string
	for i, _ := range s {
		odd := palindrome(s, i, i)
		even := palindrome(s, i, i+1)

		if len(even) > len(res) {
			res = even
		}
		if len(odd) > len(res) {
			res = odd
		}
	}
	return res
}

func palindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
