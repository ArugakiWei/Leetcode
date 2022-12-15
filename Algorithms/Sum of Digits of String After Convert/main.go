package main

import "fmt"

func main() {
	fmt.Println(getLucky("iiii", 1))
	fmt.Println(getLucky("leetcode", 2))
}

func getLucky(s string, k int) int {
	m := map[byte]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 1, 'k': 2, 'l': 3, 'm': 4,
		'n': 5, 'o': 6, 'p': 7, 'q': 8, 'r': 9, 's': 10, 't': 2, 'u': 3, 'v': 4, 'w': 5, 'x': 6, 'y': 7, 'z': 8,
	}

	ans := 0
	for _, b := range s {
		ans += m[byte(b)]
	}
	for i := 0; i < k-1; i++ {
		sum := 0
		for ans > 0 {
			sum += ans % 10
			ans = ans / 10
		}
		ans = sum
	}
	return ans
}
