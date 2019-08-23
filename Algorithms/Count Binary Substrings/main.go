package main

import (
	"fmt"
)

func countBinarySubstrings(s string) int {

	sum, pre, cur := 0, 0, 1
	for index := 1; index < len(s); index++ {
		if s[index-1] != s[index] {
			if pre < cur {
				sum += pre
			} else {
				sum += cur
			}
			pre, cur = cur, 1
		} else {
			cur++
		}
	}
	if pre < cur {
		sum += pre
	} else {
		sum += cur
	}
	return sum
}

func main() {
	a := "00110011"
	b := "10101"
	c := "00"
	d := "01100"
	e := "00100"
	f := "01"
	fmt.Println(countBinarySubstrings(a))
	fmt.Println(countBinarySubstrings(b))
	fmt.Println(countBinarySubstrings(c))
	fmt.Println(countBinarySubstrings(d))
	fmt.Println(countBinarySubstrings(e))
	fmt.Println(countBinarySubstrings(f))
}
