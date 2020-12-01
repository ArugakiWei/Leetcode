package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := Constructor([]int{1, 2, 3})
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
	fmt.Println(s.Shuffle())
}

type Solution struct {
	Original []int
	R        *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		Original: nums,
		R:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.Original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	res := make([]int, len(this.Original))
	copy(res, this.Original)

	swap := func(arr []int, i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i := 0; i < len(this.Original); i++ {
		j := this.R.Intn(len(this.Original)-i) + i
		swap(res, i, j)
	}
	return res
}
