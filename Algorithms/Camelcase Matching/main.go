package main

import "fmt"

func main() {
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBaT"))
	//fmt.Println(camelMatch([]string{"FooBar"}, "FoBaT"))
}

func camelMatch(queries []string, pattern string) []bool {
	var ans []bool
	for _, q := range queries {
		ans = append(ans, isMatch(q, pattern))
	}
	return ans
}

func isMatch(query, pattern string) bool {
	p := 0
	for i := 0; i < len(query); i++ {
		// 还没开始匹配
		if p == 0 && isUppercase(query[i]) && query[i] != pattern[p] {
			return false
		}
		// 匹配结束
		if p > len(pattern)-1 {
			if isUppercase(query[i]) {
				return false
			}
			continue
		}
		// 匹配中
		if query[i] == pattern[p] {
			p++
			continue
		}
		if isUppercase(query[i]) {
			return false
		}
	}
	return p > len(pattern)-1
}

func isUppercase(s uint8) bool {
	return s >= 'A' && s <= 'Z'
}
