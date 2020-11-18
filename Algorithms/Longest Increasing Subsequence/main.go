package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	//fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	//fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	//fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
}

/*
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	top := make([]int, len(nums))
	piles := 0

	for _, poker := range nums {
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] == poker {
				right = mid
			}
			if top[mid] > poker {
				right = mid
			}
			if top[mid] < poker {
				left = mid + 1
			}
		}

		top[left] = poker
		if left == piles {
			piles++
		}
	}

	return piles
}
