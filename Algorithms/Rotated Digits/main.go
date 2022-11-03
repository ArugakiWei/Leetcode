package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rotatedDigits(100))
}

func rotatedDigits(n int) int {
	ans := 0
	for i := 0; i <= n; i++ {
		j := i
		newSum, count, isInvalid := 0, 0, true
		for j > 0 {
			r := rotate(j % 10)
			if r == -1 {
				isInvalid = false
				break
			}
			newSum += r * int(math.Pow(10, float64(count)))
			count++
			j = j / 10
		}
		if isInvalid && newSum != i {
			ans++
		}
	}
	return ans
}

func rotate(v int) int {
	switch v {
	case 0, 1, 8:
		return v
	case 2:
		return 5
	case 5:
		return 2
	case 6:
		return 9
	case 9:
		return 6
	default:
		return -1
	}
}
