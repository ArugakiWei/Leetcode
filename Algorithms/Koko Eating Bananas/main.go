package main

import "fmt"

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}

func minEatingSpeed(piles []int, H int) int {
	left, rgiht := 1, getMax(piles)
	for left <= rgiht {
		mid := left + (rgiht-left)/2
		if canFinish(piles, mid, H) {
			rgiht = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func canFinish(piles []int, K, H int) bool {
	hours := 0
	for _, n := range piles {
		hours = hours + (n / K)
		if n%K != 0 {
			hours++
		}
	}

	return hours <= H
}

func getMax(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
