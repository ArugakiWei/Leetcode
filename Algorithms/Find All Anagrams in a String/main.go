package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

func findAnagrams(s string, p string) []int {
	needs, window := make(map[byte]int), make(map[byte]int)
	for _, b := range []byte(p) {
		needs[b]++
	}

	res := make([]int, 0)
	left, right, match := 0, 0, 0
	for right < len(s) {
		rightChar := []byte(s)[right]
		if _, ok := needs[rightChar]; ok {
			window[rightChar]++
			if window[rightChar] == needs[rightChar] {
				match++
			}
		}
		right++

		for match == len(needs) {
			if right-left == len(p) {
				res = append(res, left)
			}

			leftChar := []byte(s)[left]
			if _, ok := needs[leftChar]; ok {
				window[leftChar]--
				if window[leftChar] < needs[leftChar] {
					match--
				}
			}
			left++
		}
	}

	return res
}
