package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, len(nums))
	x := 0
	for index := 0; index < len(nums); index++ {
		res[nums[index]-1] = -1
	}
	for index := 0; index < len(res); index++ {
		if res[index] == 0 {
			x++
			res = append(res, index+1)
		}
	}
	return res[len(res)-x:]
}

func main() {
	a := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(a))
}
