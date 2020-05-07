package main

import "fmt"

func numberOfSteps (num int) int {
	switch num {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	}

	for count:=0; ; count++ {
		if num == 0 {
			return count
		}
		if num % 2 == 0 {
			num = num / 2
			continue
		}
		num = num - 1
	}
}

func numberOfSteps1 (num int) int {
	switch num {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	}

	count := 1
	if num % 2 == 0 {
		count += numberOfSteps1(num/2)
	} else {
		count += numberOfSteps1(num-1)
	}
	return count
}

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
	fmt.Println(numberOfSteps(0))
	fmt.Println(numberOfSteps(1))
	fmt.Println(numberOfSteps(2))

	fmt.Println(numberOfSteps1(14))
	fmt.Println(numberOfSteps1(8))
	fmt.Println(numberOfSteps1(123))
	fmt.Println(numberOfSteps1(0))
	fmt.Println(numberOfSteps1(1))
	fmt.Println(numberOfSteps1(2))
}