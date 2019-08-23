package main

import (
	"fmt"
	"math"
)

func main() {

	a := 5
	fmt.Println(findComplement(a))
}

func findComplement(num int) int {

	result := 0
	for i, v := range toBinaryComplementNumber(num) {
		result += int(float64(v-48) * math.Pow(2, float64(i)))
	}
	return result
}

func toBinaryComplementNumber(num int) string {

	binary := ""
	for {
		if num/2 == 0 {
			binary += "0"
			return binary
		}
		if num%2 == 0 {
			binary += "1"
		} else {
			binary += "0"
		}
		num = num / 2
	}
	return ""
}
