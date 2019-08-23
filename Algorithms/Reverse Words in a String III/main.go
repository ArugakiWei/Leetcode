package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

func reverseWords(s string) string {

	var strs []string
	for _, word := range strings.Split(s, " ") {
		rw := reverseString(word)
		strs = append(strs, rw)
	}
	return strings.Join(strs, " ")
}

func reverseString(s string) string {

	var rs []byte
	for i := len(s) - 1; i >= 0; i-- {

		rs = append(rs, s[i])
	}
	return string(rs)
}
