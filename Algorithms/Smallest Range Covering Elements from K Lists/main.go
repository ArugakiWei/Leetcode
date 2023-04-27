package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(smallestRange([][]int{
		{-5, -4, -3, -2, -1, 1},
		{1, 2, 3, 4, 5},
	}))
	fmt.Println(smallestRange([][]int{
		{4, 10, 15, 24, 26},
		{0, 9, 12, 20},
		{5, 18, 22, 30},
	}))
	fmt.Println(smallestRange([][]int{
		{1, 3, 5, 7, 9, 10},
		{2, 4, 6, 8, 10},
	}))
}

type Num struct {
	Val  int
	From int
}

func smallestRange(nums [][]int) []int {
	var vals []Num
	for i, v := range nums {
		for _, vv := range v {
			vals = append(vals, Num{
				Val:  vv,
				From: i,
			})
		}
	}

	sort.Slice(vals, func(i, j int) bool {
		if vals[i].Val < vals[j].Val {
			return true
		}
		return false
	})

	m := make(map[int]int)
	left, right := 0, 0
	start, end := 0, math.MaxInt
	for right < len(vals) {
		m[vals[right].From]++
		for len(m) == len(nums) {
			lc, la := vals[right].Val-vals[left].Val, end-start
			if lc < la || (lc == la && left < start) {
				start = vals[left].Val
				end = vals[right].Val
			}
			if _, ok := m[vals[left].From]; ok {
				m[vals[left].From]--
			}
			if m[vals[left].From] == 0 {
				delete(m, vals[left].From)
			}
			left++
		}
		right++
	}
	return []int{start, end}
}
