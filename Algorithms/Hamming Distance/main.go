package main

import "fmt"

func hammingDistance(x int, y int) int {

	var x2 []int
	var y2 []int
	i := x
	z := y
	sum := 0
	for i > 0 || z > 0 {
		if i%2 == 0 {
			x2 = append(x2, 0)
		} else {
			x2 = append(x2, 1)
		}
		if z%2 == 0 {
			y2 = append(y2, 0)
		} else {
			y2 = append(y2, 1)
		}
		i = i / 2
		z = z / 2
	}
	for i := 0; i < len(x2); i++ {
		if x2[i] != y2[i] {
			sum++
		}
	}
	return sum
}

func main() {

	fmt.Println(hammingDistance(1, 4))
}
