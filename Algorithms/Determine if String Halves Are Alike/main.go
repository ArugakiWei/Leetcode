package main

import "fmt"

func main() {
	fmt.Println(halvesAreAlike("book"))
}

func halvesAreAlike(s string) bool {
	s1 := s[:len(s)/2]
	s2 := s[len(s)/2:]
	c1, c2 := 0, 0
	for _, v := range s1 {
		if isVowels(byte(v)) {
			c1++
		}
	}
	for _, v := range s2 {
		if isVowels(byte(v)) {
			c2++
		}
	}
	return c1 == c2
}

func isVowels(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
