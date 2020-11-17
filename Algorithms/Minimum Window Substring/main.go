package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}

func minWindow(s string, t string) string {
	needs, window := make(map[byte]int), make(map[byte]int)
	for _, b := range []byte(t) {
		needs[b]++
	}

	start, minLen := 0, len(s)+1
	left, right, match := 0, 0, 0
	for right < len(s) {
		curRightChar := []byte(s)[right]
		if _, ok := needs[curRightChar]; ok {
			window[curRightChar]++
			if window[curRightChar] == needs[curRightChar] {
				match++
			}
		}
		right++

		for match == len(needs) {
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			curLeftChar := []byte(s)[left]
			if _, ok := needs[curLeftChar]; ok {
				window[curLeftChar]--
				if window[curLeftChar] < needs[curLeftChar] {
					match--
				}
			}
			left++
		}
	}

	if minLen != len(s)+1 {
		return s[start : start+minLen]
	}
	return ""
}

/*
func minWindow(s string, t string) string {
	window, needs := make(map[byte]int), make(map[byte]int)
	for _, b := range []byte(t) {
		window[b] = 0
	}
	for _, b := range []byte(t) {
		needs[b]++
	}

	left, right, ansLeft, ansRight, ansLength := 0, 0, 0, 0, len(s)+1
	for right < len(s) {
		curRightChar := []byte(s)[right]
		if _, ok := window[curRightChar]; ok {
			window[curRightChar]++
		}
		right++

		for isShrink(window, needs) {
			curLeftChar := []byte(s)[left]
			if _, ok := window[curLeftChar]; ok {
				window[curLeftChar]--
			}

			if right-left < ansLength {
				ansLeft = left
				ansRight = right
				ansLength = right - left
			}
			left++
		}
	}

	return s[ansLeft:ansRight]
}

func isShrink(window, needs map[byte]int) bool {
	for b, count := range window {
		if needsCount, ok := needs[b]; ok {
			if count < needsCount {
				return false
			}
		}
	}
	return true
}
*/
