package main

import (
	"fmt"
	"math"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	count := 0
	for _, v1 := range arr1 {
		flag := false
		for _, v2 := range arr2 {
			if math.Abs(float64(v1)-float64(v2)) <= float64(d) {
				flag = true
				break
			}
		}
		if !flag {
			count++
			flag = false
		}
	}
	return count
}

func main() {
	arr1 := []int{4,5,8}
	arr2 := []int{10,9,1,8}
	fmt.Println(findTheDistanceValue(arr1, arr2, 2))

	arr3 := []int{1,4,2,3}
	arr4 := []int{-4,-3,6,10,20,30}
	fmt.Println(findTheDistanceValue(arr3, arr4, 3))

	arr5 := []int{2,1,100,3}
	arr6 := []int{-5,-2,10,-3,7}
	fmt.Println(findTheDistanceValue(arr5, arr6, 6))
}