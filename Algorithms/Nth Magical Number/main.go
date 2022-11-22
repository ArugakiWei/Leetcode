package main

import "fmt"

func main() {
	fmt.Println(nthMagicalNumber(1, 2, 3))
	fmt.Println(nthMagicalNumber(2, 2, 3))
	fmt.Println(nthMagicalNumber(3, 2, 3))
	fmt.Println(nthMagicalNumber(4, 2, 3))
	fmt.Println(nthMagicalNumber(5, 2, 3))
	fmt.Println(nthMagicalNumber(6, 2, 3))
	fmt.Println(nthMagicalNumber(7, 2, 3))
	fmt.Println(nthMagicalNumber(8, 2, 3))
	fmt.Println(nthMagicalNumber(9, 2, 3))
}

func nthMagicalNumber(n int, a int, b int) int {
	m := min(a, b)
	c := lcm(a, b)
	left, right := m, n*m
	for left <= right {
		mid := left + (right-left)/2
		nn := nNumer(mid, a, b, c)
		if nn >= n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return mod(left)
}

func nNumer(x, a, b, c int) int {
	return x/a + x/b - x/c
}

func mod(x int) int {
	return x % 1000000007
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lcm(x, y int) int {
	return x * y / gcdx(x, y)
}

func gcdx(x, y int) int {
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
