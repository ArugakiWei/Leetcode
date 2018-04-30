package main

import "fmt"

func getSum(a int, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
	return sum
}

func main() {
	fmt.Println(getSum(3, 2))
}
