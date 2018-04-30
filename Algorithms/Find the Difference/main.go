package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	sc := make(map[byte]int, 26)
	tc := make(map[byte]int, 26)
	for i := 97 ; i < 123; i++ {
		sc[byte(i)] = 0
		tc[byte(i)] = 0
	}
	index := 0
	for ; index < len(s); index++ {
		sc[s[index]]++
		tc[t[index]]++
	}
	tc[t[index]] = tc[t[index]] + 1
	var res byte
	for i := 97; i< 123; i++ {
		if tc[byte(i)] - sc[byte(i)] == 1 {
			res = byte(i)
			break
		}
	}
	return res
}

func main() {

	s := "ab"
	t := "aab"
	fmt.Println(string(findTheDifference(s, t)))
}
