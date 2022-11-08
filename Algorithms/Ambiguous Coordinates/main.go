package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ambiguousCoordinates("(123)"))
	fmt.Println(ambiguousCoordinates("(00011)"))
	fmt.Println(ambiguousCoordinates("(0123)"))
	fmt.Println(ambiguousCoordinates("(100)"))
}

func ambiguousCoordinates(s string) []string {
	ans1 := make([]string, 0, 0)
	// 减掉2个括号, 再减最后一个y坐标
	for i := 1; i < len(s)-2; i++ {
		var a []byte
		a = append(a, '(')
		for j := 1; j < len(s)-1; j++ {
			a = append(a, s[j])
			if i == j {
				a = append(a, ',')
			}
		}
		a = append(a, ')')
		ans1 = append(ans1, string(a))
	}
	ans2 := make([]string, 0, 0)
	for _, v := range ans1 {
		ans2 = append(ans2, v)
		a := []byte(v)
		for i := 1; i < len(v)-1; i++ {
			b := make([]byte, len(a))
			copy(b, a)
			b = append(b, 0)
			copy(b[i+1:], b[i:])
			b[i] = '.'
			ans2 = append(ans2, string(b))
		}
	}
	m := make(map[string]struct{})
	for _, v := range ans2 {
		if isValid([]byte(v)) {
			m[v] = struct{}{}
		}
		a := []byte(v)
		for i := 1; i < len(v)-1; i++ {
			b := make([]byte, len(a))
			copy(b, a)
			b = append(b, 0)
			copy(b[i+1:], b[i:])
			b[i] = '.'
			if isValid(b) {
				m[string(b)] = struct{}{}
			}
		}
	}
	ans := make([]string, 0, 0)
	for v, _ := range m {
		t := make([]byte, 0, len(v))
		for _, b := range v {
			t = append(t, byte(b))
			if b == ',' {
				t = append(t, ' ')
			}
		}
		ans = append(ans, string(t))
	}
	return ans
}

func isValid(s []byte) bool {
	if strings.Contains(string(s), "..") {
		return false
	}
	if strings.Contains(string(s), ".,") || strings.Contains(string(s), ",.") {
		return false
	}
	s = s[1 : len(s)-1]
	p := strings.Split(string(s), ",")
	x, y := p[0], p[1]
	return isValidNumber(x) && isValidNumber(y)
}

func isValidNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '.' {
		return false
	}
	if strings.Count(s, ".") > 1 {
		return false
	}
	// 整数以0开头
	if !strings.Contains(s, ".") && strings.HasPrefix(s, "0") && len(s) != 1 {
		return false
	}
	if strings.Contains(s, ".") {
		// 小数以0结尾
		if strings.HasSuffix(s, "0") {
			return false
		}
		// 以0开头,但小数点前有多个0
		if strings.HasPrefix(s, "0") && strings.IndexByte(s, '.') != 1 {
			return false
		}
	}
	return true
}
