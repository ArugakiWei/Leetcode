package main

import "fmt"

func main() {
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
	fmt.Println(applyOperations([]int{1, 0}))
}

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
	}
	ans := make([]int, 0, len(nums))
	for _, v := range nums {
		if v != 0 {
			ans = append(ans, v)
		}
	}
	x := len(ans)
	for i := 0; i < len(nums)-x; i++ {
		ans = append(ans, 0)
	}
	return ans
}
