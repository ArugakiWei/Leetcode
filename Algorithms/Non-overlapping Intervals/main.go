package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
}

type Intervals [][]int

func (m Intervals) Len() int {
	return len(m)
}

func (m Intervals) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}

func (m Intervals) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(Intervals(intervals))

	x, res := intervals[0], 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < x[1] {
			res++
		} else {
			x = intervals[i]
		}
	}
	return res
}
