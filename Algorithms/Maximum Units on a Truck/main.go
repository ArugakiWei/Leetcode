package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
	fmt.Println(maximumUnits([][]int{{1, 3}, {5, 5}, {2, 5}, {4, 2}, {4, 1}, {3, 1}, {2, 2}, {1, 3}, {2, 5}, {3, 2}}, 35))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	b := Boxes(boxTypes)
	sort.Sort(b)

	max, truck := 0, 0
	for i := len(b) - 1; i >= 0; i-- {
		c := b[i][0]
		if truckSize-truck < c {
			c = truckSize - truck
		}
		truck += c
		max += c * b[i][1]
		if truck == truckSize {
			break
		}
	}
	return max
}

type Boxes [][]int

func (x Boxes) Len() int           { return len(x) }
func (x Boxes) Less(i, j int) bool { return x[i][1] < x[j][1] }
func (x Boxes) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
