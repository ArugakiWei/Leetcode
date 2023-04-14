package main

func mostFrequentEven(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if v%2 == 0 {
			m[v]++
		}
	}
	var max int
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	var ans []int
	for k, v := range m {
		if v == max {
			ans = append(ans, k)
		}
	}
	if len(ans) == 0 {
		return -1
	}
	min := 100001
	for _, v := range ans {
		if v < min {
			min = v
		}
	}
	return min
}
