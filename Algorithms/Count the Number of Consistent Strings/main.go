package main

import "fmt"

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
}

func countConsistentStrings(allowed string, words []string) int {
	m := make(map[int32]struct{})
	for _, v := range allowed {
		m[v] = struct{}{}
	}
	ans := len(words)
	for _, w := range words {
		for _, v := range w {
			if _, ok := m[v]; !ok {
				ans--
				break
			}
		}
	}
	return ans
}
