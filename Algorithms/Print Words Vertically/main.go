package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(printVertically("TO BE OR NOT TO BE"))
	// fmt.Println(printVertically("TO BE OR NO TO BE"))
}

func printVertically(s string) []string {
	ss := strings.Split(s, " ")
	max := len(ss[0])
	for _, v := range ss {
		if len(v) > max {
			max = len(v)
		}
	}

	result := make([]string, max)
	for i := 0; i < max; i++ {
		ri := make([]byte, len(ss))
		for j := 0; j < len(ss); j++ {
			if i > len(ss[j])-1 {
				ri[j] = ' '
			} else {
				ri[j] = ss[j][i]
			}
		}
		result[i] = strings.TrimRight(string(ri), " ")
	}
	return result
}
