package main

import "fmt"

func main() {
	fmt.Println(digitCount("030"))
}

func digitCount(num string) bool {
	m := make(map[int]int)
	for _, v := range num {
		vv := v - '0'
		m[int(vv)]++
	}
	for i, v := range num {
		vv := v - '0'
		if int(vv) != m[i] {
			return false
		}
	}
	return true
}
