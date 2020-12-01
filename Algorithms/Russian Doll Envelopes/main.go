package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxEnvelopes([][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}, {1, 1}}))
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Sort(Envelopes(envelopes))

	h := make([]int, len(envelopes))
	for i, v := range envelopes {
		h[i] = v[1]
	}

	return lengthOfLIS(h)
}

type Envelopes [][]int

func (m Envelopes) Len() int {
	return len(m)
}

func (m Envelopes) Less(i, j int) bool {
	if m[i][0] < m[j][0] {
		return true
	}
	if m[i][0] > m[j][0] {
		return false
	}
	if m[i][0] == m[j][0] && m[i][1] > m[j][1] {
		return true
	}
	return false
}

func (m Envelopes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	top := make([]int, len(nums))
	piles := 0

	for _, poker := range nums {
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)/2
			if top[mid] == poker {
				right = mid
			}
			if top[mid] > poker {
				right = mid
			}
			if top[mid] < poker {
				left = mid + 1
			}
		}

		top[left] = poker
		if left == piles {
			piles++
		}
	}

	return piles
}
