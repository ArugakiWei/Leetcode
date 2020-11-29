package main

import "fmt"

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
	fmt.Println(pancakeSort([]int{1, 2, 3}))
}

func pancakeSort(arr []int) []int {
	var res []int
	var sort func(arr []int, n int)
	sort = func(arr []int, n int) {
		if n == 1 {
			return
		}

		max := 0
		for i := 0; i < n; i++ {
			if arr[max] < arr[i] {
				max = i
			}
		}

		reserve(arr, 0, max)
		res = append(res, max+1)
		reserve(arr, 0, n-1)
		res = append(res, n)

		sort(arr, n-1)
	}

	sort(arr, len(arr))
	return res
}

func reserve(arr []int, top, bottom int) {
	for i, j := top, bottom; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
