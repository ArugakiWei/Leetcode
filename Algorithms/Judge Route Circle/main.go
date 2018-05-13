package main

import "fmt"

func main() {

	a := "LLLRRR"
	fmt.Println(judgeCircle(a))
}

func judgeCircle(moves string) bool {

	if len(moves) % 2 != 0 {
		return false
	}
	u, d, l, r := 0, 0, 0, 0
	for _, c := range moves {
		if c == 'U' {
			u++
		}
		if c == 'D' {
			d++
		}
		if c == 'L' {
			l++
		}
		if c == 'R' {
			r++
		}
	}
	if u == d && l == r {
		return true
	}
	return false
}
