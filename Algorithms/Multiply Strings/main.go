package main

import "fmt"

func main() {
	fmt.Println(multiply("0", "0"))
	fmt.Println(multiply("103", "456"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := make([]uint8, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			nn := num1[i] - '0'
			jn := num2[j] - '0'

			if nn*jn >= 10 {
				res[i+j+1] += nn * jn % 10
				res[i+j] += nn * jn / 10
			} else {
				res[i+j+1] += nn * jn
			}

			if res[i+j+1] >= 10 {
				res[i+j+1] -= 10
				res[i+j] += 1
			}
			if res[i+j] >= 10 {
				res[i+j] -= 10
				res[i+j-1] += 1
			}
		}
	}

	i := 0
	for ; i < len(res); i++ {
		if res[i] != 0 {
			break
		}
	}
	for j := i; j < len(res); j++ {
		res[j] += '0'
	}
	return string(res[i:])
}
