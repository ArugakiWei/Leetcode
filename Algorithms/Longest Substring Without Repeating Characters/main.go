package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

/*
// Brute Force
func lengthOfLongestSubstring(s string) int {
	allUnique := func(s string, start, end int) bool {
		set := make(map[uint8]bool)
		for i:=start; i<end; i++ {
			if _, ok := set[s[i]]; ok {
				return false
			}
			set[s[i]] = true
		}
		return true
	}

	var ans int
	for i:=0; i<len(s); i++ {
		for j:=i+1; j<=len(s); j++ {
			if allUnique(s, i, j) {
				ans = int(math.Max(float64(ans), float64(j-i)))
			}
		}
	}
	return ans
}
*/

/*
// Sliding Window
func lengthOfLongestSubstring(s string) int {
	set := make(map[uint8]bool)
	ans, i, j := 0, 0, 0

	for i<len(s) && j<len(s) {
		if _, ok := set[s[j]]; !ok {
			set[s[j]] = true
			j++
			ans = int(math.Max(float64(ans), float64(j-i)))
		} else {
			delete(set, s[i])
			i++
		}
	}

	return ans
}
*/

/*
func lengthOfLongestSubstring(s string) int {
	set := make(map[uint8]bool)
	ans, i, j := 0, 0, 0

	for i<len(s) && j<len(s) {
		if _, ok := set[s[j]]; !ok {
			set[s[j]] = true
			j++
			ans = int(math.Max(float64(ans), float64(j-i)))
		} else {
			delete(set, s[i])
			i++
		}
	}

	return ans
}
*/

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)

	left, right, maxLen := 0, 0, 0
	for right < len(s) {
		rightChar := []byte(s)[right]
		window[rightChar]++
		right++

		for window[rightChar] > 1 {
			leftChar := []byte(s)[left]
			window[leftChar]--
			left++
		}

		maxLen = max(maxLen, right-left)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
