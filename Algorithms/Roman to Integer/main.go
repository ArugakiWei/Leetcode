package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if s[i] == 'I' && s[i+1] == 'V' {
				res += 4
				i++
				continue
			}
			if s[i] == 'I' && s[i+1] == 'X' {
				res += 9
				i++
				continue
			}
			if s[i] == 'X' && s[i+1] == 'L' {
				res += 40
				i++
				continue
			}
			if s[i] == 'X' && s[i+1] == 'C' {
				res += 90
				i++
				continue
			}
			if s[i] == 'C' && s[i+1] == 'D' {
				res += 400
				i++
				continue
			}
			if s[i] == 'C' && s[i+1] == 'M' {
				res += 900
				i++
				continue
			}
		}
		res += dict[s[i]]
	}
	return res
}
