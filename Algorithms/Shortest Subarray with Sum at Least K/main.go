package main

import "fmt"

func main() {
	fmt.Println(shortestSubarray([]int{2, -1, 2}, 3))
	fmt.Println(shortestSubarray([]int{1, 2}, 4))
	fmt.Println(shortestSubarray([]int{1}, 1))
}

func shortestSubarray(nums []int, k int) int {
	ans := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		curLen, curSum := 1, nums[i]
		for j := i + 1; j < len(nums); j++ {
			if curSum >= k {
				if curLen < ans {
					ans = curLen
				}
				break
			}
			curSum += nums[j]
			curLen++
		}
		if curSum >= k {
			if curLen < ans {
				ans = curLen
			}
		}
	}
	if ans == len(nums)+1 {
		return -1
	}
	return ans
}
