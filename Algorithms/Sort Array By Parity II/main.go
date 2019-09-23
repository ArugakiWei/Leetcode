package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	result := make([]int, len(A), len(A))

	for i, odd, even := 0, 0, 1; i < len(A); i++ {
		if A[i] % 2 == 0 {
			result[odd] = A[i]
			odd = odd + 2
		} else {
			result[even] = A[i]
			even = even + 2
		}
	}

	return result
}

func main() {

	a := []int{4,2,3,1}
	fmt.Println(sortArrayByParityII(a))
}