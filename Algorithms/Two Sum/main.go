package main

import (
	"fmt"
)

// O(n^2)
func twoSum(nums []int, target int) []int {
	var result []int
	for i1, v1 := range nums {
		for i2:=i1+1; i2<len(nums); i2++ {
			if v1 + nums[i2] == target {
				result = append(result, i1, i2)
				return result
			}
		}
	}
	return result
}

// O(n)
func twoSum1(nums []int, target int) []int {
	var result []int

	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}

	for i1, v1 := range nums {
		v2 := target - v1
		if is, ok := m[v2]; ok {
			if len(is) == 1 && is[0] != i1 {
				return append(result, i1, is[0])
			}
			if len(is) > 1 {
				for _, i3 := range is {
					if i3 != i1 {
						return append(result, i1, i3)
					}
				}
			}
		}
	}
	return result
}


func main() {
	a1 := []int{2, 7, 11, 15}
	fmt.Println(twoSum1(a1, 9))

	a2 := []int{1, 4, 3, 12}
	fmt.Println(twoSum1(a2, 15))

	a3 := []int{3, 2, 4}
	fmt.Println(twoSum1(a3, 6))

	a4 := []int{3, 3}
	fmt.Println(twoSum1(a4, 6))
}