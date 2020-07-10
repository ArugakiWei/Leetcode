package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("bbbbb"))
	fmt.Println(longestPalindrome("abcda"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1{
		return s
	}

	m := make(map[string]int)
	for i:=0; i<len(s)-1; i++ {
		for j:=i+1; j<len(s); j++ {
			if s[j] == s[i] {
				ri, ti := i, i+1
				rj, tj := j, j-1
				var flag bool
				for ti < tj {
					if s[ti] != s[tj] {
						flag = true
						break
					}
					ti++
					tj--
				}
				if !flag {
					m[fmt.Sprintf("%d-%d", ri, rj)] = rj-ri
				}
			}
		}
	}

	var max int
	var start, end string
	for k, v := range m {
		if v > max {
			max = v
			start = strings.Split(k, "-")[0]
			end = strings.Split(k, "-")[1]
		}
	}

	i, _ := strconv.ParseInt(start, 10, 64)
	j, _ := strconv.ParseInt(end, 10, 64)
	return s[i:j+1]
}