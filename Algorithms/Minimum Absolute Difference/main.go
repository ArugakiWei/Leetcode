package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4,2,1,3}))
	fmt.Println(minimumAbsDifference([]int{1,3,6,10,15}))
	fmt.Println(minimumAbsDifference([]int{3,8,-10,23,19,-4,-14,27}))
	fmt.Println(minimumAbsDifference([]int{40,11,26,27,-20}))
}

//func minimumAbsDifference(arr []int) [][]int {
//	arr = sort(arr)
//
//	min := arr[1] - arr[0]
//	pairs := make([][]int, 0, 0)
//	for i := 0; i < len(arr) - 1; i++ {
//		distinct := arr[i+1] - arr[i]
//		if distinct < min {
//			min = distinct
//			newPairs := make([][]int, 0, 0)
//			pair := make([]int, 0, 2)
//			pair = append(pair, arr[i], arr[i+1])
//			newPairs = append(newPairs, pair)
//			pairs = newPairs
//		} else if distinct == min {
//			pair := make([]int, 0, 2)
//			pair = append(pair, arr[i], arr[i+1])
//			pairs = append(pairs, pair)
//		}
//	}
//	return pairs
//}
//
//func sort(arr []int) []int {
//	max := arr[0]
//	for _, v := range arr {
//		if v < 0 {
//			v = -v
//		}
//		if v > max {
//			max = v
//		}
//	}
//	arrLen := max * 2 + 1
//	newArr := make([]int, arrLen, arrLen)
//	middle := arrLen / 2
//	for _, v := range arr {
//		newArr[middle+v]++
//	}
//	result := make([]int, 0, len(arr))
//	for i, v := range newArr {
//		if v != 0 {
//			result = append(result, i - middle)
//		}
//	}
//	return result
//}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	min := arr[1] - arr[0]
	pairs := make([][]int, 0, 0)
	for i := 0; i < len(arr) - 1; i++ {
		distinct := arr[i+1] - arr[i]
		if distinct <= min {
			if distinct < min {
				min = distinct
				pairs = make([][]int, 0, 0)
			}
			pair := make([]int, 0, 2)
			pair = append(pair, arr[i], arr[i+1])
			pairs = append(pairs, pair)
		}
	}
	return pairs
}