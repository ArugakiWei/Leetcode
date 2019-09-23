package main

import (
	"fmt"
	"math"
)

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}

	a := math.Pow(5, 0.5)
	b := math.Pow((1+a)/2, float64(N))
	c := math.Pow((1-a)/2, float64(N))
	fn := ((1/a) * (b-c)) + 0.5
	return int(math.Floor(fn))
}

func main() {
	fmt.Println(fib(12))
}
