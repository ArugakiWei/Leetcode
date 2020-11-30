package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{2, 6}, {1, 3}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}

func merge(intervals [][]int) [][]int {
	sort.Sort(Intervals(intervals))

	res := make([][]int, 0)
	t := intervals[0]
	for i := 1; i < len(intervals); i++ {
		// intervals[i] 是 t 的子区间
		if intervals[i][0] <= t[1] && intervals[i][1] <= t[1] {
			continue
		}
		// intervals[i] 的 end 值更大, 与 t 合并
		if intervals[i][0] <= t[1] && intervals[i][1] >= t[1] {
			t[1] = intervals[i][1]
		}
		// 没有交集
		if intervals[i][0] > t[1] {
			res = append(res, t)
			t = intervals[i]
		}
	}
	res = append(res, t)
	return res
}

type Intervals [][]int

func (m Intervals) Len() int {
	return len(m)
}

func (m Intervals) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

func (m Intervals) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
