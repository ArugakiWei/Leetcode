package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {

	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	} else {
		return num % 9
	}
}
