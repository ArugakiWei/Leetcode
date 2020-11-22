package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}

func isMatch(s string, p string) bool {
	memo := make(map[string]bool)
	return dp(s, p, 0, 0, memo)
}

func dp(s, p string, i, j int, memo map[string]bool) bool {
	// p已经匹配完, 判断s是否匹配完
	if j == len(p) {
		return i == len(s)
	}
	// i已经匹配完
	if i == len(s) {
		// 如果能匹配空串，一定是字符和 * 成对儿出现
		if (len(p) - j) % 2 == 1 {
			return false
		}
		// 检查是否为 x*y*z* 这种形式
		for ; j < len(p) - 1; j+=2 {
			if p[j+1] != '*' {
				return false
			}
		}
		return true
	}

	key := fmt.Sprintf("%d%d", i, j)
	if v, ok := memo[key]; ok {
		return v
	}

	var res bool
	if s[i] == p[j] || p[j] == '.' {
		// 下个字符是 * , 可以匹配0次或多次
		if j < len(p) - 1 && p[j+1] == '*' {
			// 匹配0次 || 匹配多次
			res = dp(s, p, i, j+2, memo) || dp(s, p, i+1, j, memo)
		} else {
			res = dp(s, p, i+1, j+1, memo)
		}
	} else {
		// 下个字符是 * , 只能匹配0次
		if j < len(p) - 1 && p[j+1] == '*' {
			res = dp(s, p, i, j+2, memo)
		}
	}
	memo[key] = res

	return res
}

