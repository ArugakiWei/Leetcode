package main

import "fmt"

func main() {
	// fmt.Println(isGoodArray([]int{12, 5, 7, 23}))
	fmt.Println(isGoodArray([]int{6, 10, 15}))
	// fmt.Println(isGoodArray([]int{3, 6}))
	// fmt.Println(isGoodArray([]int{12, 30, 6, 54, 42, 48, 12, 18, 48, 54, 90, 50, 50, 70, 40, 90, 100, 100, 10, 90, 5}))
}

func isGoodArray(nums []int) bool {
	x := 0
	for _, n := range nums {
		x = gcd(x, n)
		if x == 1 {
			return true
		}
	}
	return false
}

func gcd(x, y int) int {
	var tmp int
	for {
		tmp = x % y
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}
