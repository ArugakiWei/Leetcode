package main

import "fmt"

func main() {

	fmt.Println(numJewelsInStones("z", "ZZZ"))
}

func numJewelsInStones(J string, S string) int {
	sum := 0
	t := make(map[int32]int)
	for _, value := range S {
		t[value]++
	}
	for _, value := range J {
		sum = sum + t[value]
	}
	return sum
}
