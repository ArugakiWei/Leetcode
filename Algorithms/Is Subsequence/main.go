package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "assdsdsdbjsadjc"))
}

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for ; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}
