package main

import "fmt"

func main() {
	// fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
	// fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
	// fmt.Println(numSubarrayBoundedMax([]int{73, 55, 36, 5, 55, 14, 9, 7, 72, 52}, 32, 69))
	// fmt.Println(numSubarrayBoundedMax([]int{16, 69, 88, 85, 79, 87, 37, 33, 39, 34}, 55, 57))
	fmt.Println(numSubarrayBoundedMax([]int{1, 7, 0, 0, 9, 3, 0, 7, 4, 2, 1, 3, 9, 0, 3, 0}, 7, 8))
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	ans := 0
	lessLen, curLen := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < left {
			lessLen++
		}
		if nums[i] >= left {
			ans -= subarraysLen(lessLen)
			lessLen = 0
		}
		if nums[i] > right {
			ans += subarraysLen(curLen)
			curLen = 0
		} else {
			curLen++
		}
	}
	ans += subarraysLen(curLen) - subarraysLen(lessLen)
	return ans
}

func subarraysLen(n int) int {
	if n%2 == 0 {
		return (n + 1) * (n / 2)
	}
	return n * (n/2 + 1)
}
