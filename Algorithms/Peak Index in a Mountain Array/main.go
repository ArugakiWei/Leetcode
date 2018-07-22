package main

import (
	"fmt"
)

func main() {

	c := []int{18,29,38,59,98,100,99,98,90}
	fmt.Println(peakIndexInMountainArray(c))
}

func peakIndexInMountainArray(A []int) int {

	start, end := 0, len(A)
	for  {
		i := start + (end-start)/2
		if A[i] > A[i-1] && A[i] > A[i+1] {
			return i
		}
		if A[i] > A[i-1] && A[i] < A[i+1] {
			start = i
			continue
		}
		if A[i] < A[i-1] && A[i] > A[i+1] {
			end = i
			continue
		}
	}
}
