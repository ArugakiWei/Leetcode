package main

import (
	"fmt"
)

func main() {

	a := []int{0, 1, 0, 3, 12}
	b := []int{0,0,1}
	c := []int{1,0,0}
	moveZeroes(a)
	moveZeroes(b)
	moveZeroes(c)
}

func moveZeroes(nums []int)  {

	j := 0
	for i:=0; j<len(nums); i++ {
		j++
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			i--
		}
	}
	fmt.Println(nums)
}
