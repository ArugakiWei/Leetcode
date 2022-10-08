package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{
		"flower", "flow", "flight",
	}))
	fmt.Println(longestCommonPrefix([]string{
		"dog", "racecar", "car",
	}))
	fmt.Println(longestCommonPrefix([]string{
		"a",
	}))
	fmt.Println(longestCommonPrefix([]string{
		"ab", "a",
	}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLenStr := strs[0]
	for _, s := range strs {
		if len(s) < len(minLenStr) {
			minLenStr = s
		}
	}

	var i int
	for i = 0; i < len(minLenStr); i++ {
		flag := true
		for _, s := range strs {
			if s[i] != minLenStr[i] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
	}
	return minLenStr[:i]
}
