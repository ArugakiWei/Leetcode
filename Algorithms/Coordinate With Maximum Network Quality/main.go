package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(bestCoordinate([][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, 2))
}

func bestCoordinate(towers [][]int, radius int) []int {
	maxX, maxY := 0, 0
	for _, t := range towers {
		if t[0] > maxX {
			maxX = t[0]
		}
		if t[1] > maxY {
			maxY = t[1]
		}
	}
	maxQ, ans := 0, []int{0, 0}
	for i := 0; i <= maxX+radius; i++ {
		for j := 0; j <= maxY+radius; j++ {
			sumQ := 0
			for _, t := range towers {
				a := []int{i, j}
				b := []int{t[0], t[1]}
				d := dis(a, b)
				if d > float64(radius) {
					continue
				}
				sumQ += signal(d, t[2])
			}
			if sumQ > maxQ {
				maxQ = sumQ
				ans[0], ans[1] = i, j
			}
		}
	}
	return ans
}

func signal(d float64, q int) int {
	return int(math.Floor(float64(q) / (1 + d)))
}

func dis(a, b []int) float64 {
	dx := abs(a[0] - b[0])
	dy := abs(a[1] - b[1])
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
