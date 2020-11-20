package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("", "a"))
	fmt.Println(strStr("aa", ""))
}

func strStr(haystack string, needle string) int {
	k := NewKMP(needle)
	return k.search(haystack)
}

type KMP struct {
	Pat string
	DP  [][]int
}

func NewKMP(pat string) *KMP {
	if len(pat) == 0 {
		return &KMP{}
	}

	dp := make([][]int, len(pat))
	for i, _ := range dp {
		dp[i] = make([]int, 256)
	}

	// 只有与第一个字符相等才推进到第一个状态
	dp[0][pat[0]] = 1

	x := 0
	for j := 1; j < len(pat); j++ {
		for c := 0; c < 256; c++ {
			if pat[j] == byte(c) {
				dp[j][c] = j + 1
			} else {
				dp[j][c] = dp[x][c]
			}
		}
		x = dp[x][pat[j]]
	}

	return &KMP{
		Pat: pat,
		DP:  dp,
	}
}

func (k KMP) search(txt string) int {
	if len(k.Pat) == 0 {
		return 0
	}

	m, j := len(k.Pat), 0
	for i, c := range txt {
		j = k.DP[j][c]

		if j == m {
			return i - m + 1
		}
	}
	return -1
}
