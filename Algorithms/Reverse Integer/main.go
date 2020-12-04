package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	toArray := func(x int) []int {
		nums := make([]int, 0, 32)
		for x != 0 {
			nums = append(nums, x%10)
			x = x / 10
		}
		return nums
	}

	nums, res := toArray(x), 0
	for i, num := range nums {
		res += num * int(math.Pow(10, float64(len(nums)-1-i)))
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
