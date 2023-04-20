package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
}

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	fmt.Println(satisfaction)
	sum, i := 0, len(satisfaction)-1
	for ; i >= 0; i-- {
		x := sum + satisfaction[i]
		if x < 0 {
			break
		} else {
			sum = x
		}
	}
	ans := 0
	i++
	for x := 1; i < len(satisfaction); i, x = i+1, x+1 {
		fmt.Println(satisfaction[i], x)
		ans += satisfaction[i] * x
	}
	return ans
}
