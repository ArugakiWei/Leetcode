package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	var res [][]int
	var backtrack func(n, k, start int, nums []int)
	backtrack = func(n, k, start int, nums []int) {
		if len(nums) == k {
			t := make([]int, len(nums))
			copy(t, nums)
			res = append(res, t)
			return
		}

		for i := start; i <= n; i++ {
			if contains(nums, i) {
				continue
			}

			nums = append(nums, i)
			backtrack(n, k, i+1, nums)
			nums = nums[:len(nums)-1]
		}
	}

	backtrack(n, k, 1, []int{})
	return res
}

func contains(nums []int, i int) bool {
	for _, v := range nums {
		if v == i {
			return true
		}
	}
	return false
}
