package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	num []int
	R   *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		num: nums,
		R:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) Pick(target int) int {
	res, count := -1, 0
	for i, v := range this.num {
		if v == target {
			count++
			if this.R.Intn(count) == 0 {
				res = i
			}
		}
	}
	return res
}
