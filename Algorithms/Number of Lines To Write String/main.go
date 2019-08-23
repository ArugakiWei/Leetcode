package main

import "fmt"

func main() {

	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))
}

func numberOfLines(widths []int, S string) []int {

	lineSum := 0
	lines := 1
	lastLine := 0
	for i := 0; i < len(S); i++ {
		lineSum += widths[S[i]-'a']
		if lineSum < 100 {
			lastLine = lineSum
			continue
		}
		if lineSum == 100 {
			lines++
			lineSum = 0
		}
		if lineSum > 100 {
			lines++
			i--
			lineSum = 0
		}
	}
	return []int{lines, lastLine}
}
