package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}

func checkInclusion(s1 string, s2 string) bool {
	needs, window := make(map[byte]int), make(map[byte]int)
	for _, b := range []byte(s1) {
		needs[b]++
	}

	left, right, match := 0, 0, 0
	for right < len(s2) {
		rightChar := []byte(s2)[right]
		window[rightChar]++
		right++

		if count, ok := needs[rightChar]; ok {
			if count == window[rightChar] {
				match++
			}
		}
		for match == len(needs) {
			if right-left == len(s1) {
				return true
			}

			leftChar := []byte(s2)[left]
			window[leftChar]--
			left++

			if count, ok := needs[leftChar]; ok {
				if window[leftChar] < count {
					match--
				}
			}
		}
	}

	return false
}
