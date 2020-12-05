package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(myAtoi("42"))
	//fmt.Println(myAtoi("    -42"))
	//fmt.Println(myAtoi("4193 with words"))
	//fmt.Println(myAtoi("words and 987"))
	//fmt.Println(myAtoi("-91283472332"))
	//fmt.Println(myAtoi("21474836460"))
	//fmt.Println(myAtoi("-000000000000001"))
	//fmt.Println(myAtoi("     +004500"))
	//fmt.Println(myAtoi("   +0 123"))
	fmt.Println(myAtoi("9223372036854775808"))
}

func myAtoi(s string) int {
	isNegative := false
	num := make([]int, 0, 32)

	// x 为有效数字的长度
	for i, x := 0, 0; i < len(s); i++ {
		if x == 0 {
			// 遇见空格跳过
			if s[i] == ' ' {
				continue
			}
			// 第一个有效字符不是数字或者+-
			if (s[i] < '0' || s[i] > '9') && s[i] != '-' && s[i] != '+' {
				return 0
			}
			// 第一个有效字符是 -
			if s[i] == '-' {
				isNegative = true
			}
			// 第一个有效字符是数字
			if s[i] >= '0' && s[i] <= '9' {
				num = append(num, int(s[i]-'0'))
			}
			x++
			continue
		}
		// 碰到第一个非法字符结束
		if s[i] < '0' || s[i] > '9' {
			break
		}
		num = append(num, int(s[i]-'0'))
		x++
	}

	res := 0
	for i := 0; i < len(num); i++ {
		res = res*10 + num[i]
		if res > math.MaxInt32 {
			if isNegative {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	if isNegative {
		return -res
	}
	return res
}
