package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPowersOfThree(12))
	fmt.Println(checkPowersOfThree(91))
	fmt.Println(checkPowersOfThree(21))
}

func checkPowersOfThree(n int) bool {
	var nums []int64
	for i := 0; i <= 17; i++ {
		x := int64(math.Pow(3, float64(i)))
		nums = append(nums, x)
		if int(x) > n {
			break
		}
	}

	subs := [][]int64{{}}
	for i := 0; i < len(nums); i++ {
		for _, sub := range subs {
			// 原子集 + 新元素
			t := make([]int64, len(sub))
			copy(t, sub)
			t = append(t, nums[i])
			subs = append(subs, t)
			if sum(t) == int64(n) {
				return true
			}
		}
	}
	return false
}

func sum(nums []int64) int64 {
	val := int64(0)
	for _, v := range nums {
		val += v
	}
	return val
}

func checkPowersOfThree1(n int) bool {
	for ; n > 0; n /= 3 {
		if n%3 == 2 {
			return false
		}
	}
	return true
}
