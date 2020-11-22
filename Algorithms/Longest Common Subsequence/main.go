package main

import "fmt"

func main() {
	//fmt.Println(longestCommonSubsequence("abcde", "ace"))
	//fmt.Println(longestCommonSubsequence("abc", "abc"))
	//fmt.Println(longestCommonSubsequence("abc", "def"))
	//fmt.Println(longestCommonSubsequence("ezupkr","ubmrapg"))
	fmt.Println(longestCommonSubsequence("abcba", "abcbcba"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(text2)+1)
		dp[i][0] = 0
	}

	// 时间复杂度: O(mn), 空间复杂度 O(mn)
	for i:=1; i<=len(text1); i++ {
		for j:=1; j<=len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]

	// 时间复杂度: O(mn), 空间复杂度 O(min(m,n))
	//j1, i1J1, res := 0, 0, 0
	//i1 := make([]int, len(text2)+1)
	//for i:=1; i<=len(text1); i++ {
	//	j1 = 0
	//	i1J1 = 0
	//	for j:=1; j<=len(text2); j++ {
	//		if text1[i-1] == text2[j-1] {
	//			res = i1J1 + 1
	//		} else {
	//			res = max(i1[j], j1)
	//		}
	//		j1 = res
	//		i1J1 = i1[j]
	//		i1[j] = res
	//	}
	//}

	//return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}