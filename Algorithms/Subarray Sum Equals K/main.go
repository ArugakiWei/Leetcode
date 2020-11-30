package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1}, 0))
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}

func subarraySum(nums []int, k int) int {
	preSumCount := make(map[int]int)
	preSumCount[0] = 1

	sumI, sumJ, ans := 0, 0, 0
	for _, v := range nums {
		sumI += v
		sumJ = sumI - k

		if c, ok := preSumCount[sumJ]; ok {
			ans += c
		}
		preSumCount[sumI]++
	}
	return ans
}
