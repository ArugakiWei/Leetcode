package main

import "fmt"

func main() {

	fmt.Println(reverseString("hello"))
}

func reverseString(s string) string {

	var rs []byte
	for i:=len(s)-1; i>=0; i-- {

		rs = append(rs, s[i])
	}
	return string(rs)
}
