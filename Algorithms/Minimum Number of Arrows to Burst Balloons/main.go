package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Sort(Intervals(points))

	x, res := points[0], 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > x[1] {
			x = points[i]
			res++
		}
	}
	return res
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
