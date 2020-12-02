package main

import "fmt"

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}

func shipWithinDays(weights []int, D int) int {
	left, right := getMax(weights), getSum(weights)
	for left <= right {
		mid := left + (right-left)/2
		if canFinish(weights, mid, D) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func canFinish(weights []int, cap, D int) bool {
	d, sum := 1, 0
	for _, w := range weights {
		if sum+w > cap {
			sum = w
			d++
		} else {
			sum += w
		}
	}
	return d <= D
}

func getMax(weights []int) int {
	max := weights[0]
	for _, v := range weights {
		if v > max {
			max = v
		}
	}
	return max
}

func getSum(weights []int) int {
	sum := 0
	for _, v := range weights {
		sum += v
	}
	return sum
}
