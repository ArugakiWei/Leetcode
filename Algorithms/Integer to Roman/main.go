package main

import "fmt"

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	symbol := [][]string{{"I", "V"}, {"X", "L"}, {"C", "D"}, {"M", ""}}

	res := ""
	count, cur := 0, 0
	for num > 0 {
		cur = num % 10
		num /= 10

		s := symbol[count]
		switch cur {
		case 1:
			res = s[0] + res
		case 2:
			res = s[0] + s[0] + res
		case 3:
			res = s[0] + s[0] + s[0] + res
		case 4:
			res = s[0] + s[1] + res
		case 5:
			res = s[1] + res
		case 6:
			res = s[1] + s[0] + res
		case 7:
			res = s[1] + s[0] + s[0] + res
		case 8:
			res = s[1] + s[0] + s[0] + s[0] + res
		case 9:
			s1 := symbol[count+1]
			res = s[0] + s1[0] + res
		}

		count++
	}
	return res
}
