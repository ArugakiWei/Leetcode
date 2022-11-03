package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maxRepeating("ababc", "ab"))
}

func maxRepeating(sequence string, word string) int {
	sumWord := word
	for i := 0; ; i++ {
		if !strings.Contains(sequence, sumWord) {
			return i
		}
		sumWord += word
	}
}
