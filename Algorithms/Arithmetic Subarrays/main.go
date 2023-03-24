package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
	fmt.Println(checkArithmeticSubarrays([]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}))
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var ans []bool
	for i := 0; i < len(l); i++ {
		dst := make([]int, r[i]+1-l[i])
		if r[i]+1 >= len(l) {
			copy(dst, nums[l[i]:])
		} else {
			copy(dst, nums[l[i]:r[i]+1])
		}
		ans = append(ans, isArithmeticSubarrays(dst))
	}
	return ans
}

func isArithmeticSubarrays(n []int) bool {
	sort.Ints(n)
	d := n[1] - n[0]
	for i := 2; i < len(n); i++ {
		if n[i]-n[i-1] != d {
			return false
		}
	}
	return true
}
