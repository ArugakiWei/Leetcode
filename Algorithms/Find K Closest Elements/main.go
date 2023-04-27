package main

import "fmt"

func main() {
	fmt.Println(findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5))
	fmt.Println(findClosestElements([]int{1, 1, 1, 10, 10, 10}, 1, 9))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 1))
}

func findClosestElements(arr []int, k int, x int) []int {
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}
	if x <= arr[0] {
		return arr[:k]
	}

	mid, left, right := 0, 0, len(arr)-1
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] == x {
			break
		}
		if arr[mid] > x {
			right--
		} else {
			left = mid + 1
		}
	}
	i, j := mid-1, mid
	for j-i-1 < k {
		if i < 0 {
			j++
			continue
		}
		if j > len(arr)-1 {
			i--
			continue
		}
		di, dj := abs(arr[i], x), abs(arr[j], x)
		if di <= dj {
			i--
		}
		if di > dj {
			j++
		}
	}
	return arr[i+1 : j]
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
