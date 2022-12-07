package main

import "fmt"

func main() {
	// fmt.Println(minOperations("0100"))
	// fmt.Println(minOperations("10"))
	fmt.Println(minOperations("10010100"))
}

func minOperations(s string) int {
	cl, cc := 0, 0
	cLast, cCurrent := []byte(s), []byte(s)
	for i := 1; i < len(s); i++ {
		if cLast[i] == cLast[i-1] {
			cLast[i-1] = diff(cLast[i])
			cl++
		}
		if cCurrent[i] == cCurrent[i-1] {
			cCurrent[i] = diff(cCurrent[i-1])
			cc++
		}
	}
	fmt.Println(string(cLast), string(cCurrent), cl, cc)
	if cl > cc {
		return cc
	}
	return cl
}

func diff(x byte) byte {
	if x == '1' {
		return '0'
	}
	return '1'
}
