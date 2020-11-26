package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

//func subsets(nums []int) [][]int {
//	if len(nums) == 0 {
//		return [][]int{}
//	}
//
//	dp := make([][][]int, len(nums))
//	dp[0] = [][]int{{}, {nums[0]}}
//
//	for i := 1; i < len(nums); i++ {
//		for _, sub := range dp[i-1] {
//			// 原子集
//			dp[i] = append(dp[i], sub)
//
//			// 原子集 + 新元素
//			t := make([]int, len(sub))
//			copy(t, sub)
//			t = append(t, nums[i])
//			dp[i] = append(dp[i], t)
//		}
//	}
//
//	return dp[len(nums)-1]
//}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	ans := [][]int{{}, {nums[0]}}

	for i := 1; i < len(nums); i++ {
		for _, sub := range ans {
			// 原子集 + 新元素
			t := make([]int, len(sub))
			copy(t, sub)
			t = append(t, nums[i])
			ans = append(ans, t)
		}
	}

	return ans
}
