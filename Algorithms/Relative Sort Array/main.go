package main

import (
	"container/heap"
	"fmt"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	eleCount := make(map[int]int)
	for _, ele := range arr2 {
		eleCount[ele] = 1
	}

	absent := &IntHeap{}
	heap.Init(absent)
	for _, ele := range arr1 {
		if _, ok := eleCount[ele]; ok {
			eleCount[ele]++
		} else {
			heap.Push(absent, ele)
		}
	}

	result := make([]int, 0, len(arr1))
	for _, ele := range arr2 {
		count := eleCount[ele]
		for i := 1; i < count; i++ {
			result = append(result, ele)
		}
	}

	for absent.Len() > 0 {
		result = append(result, heap.Pop(absent).(int))
	}
	return result
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	arr1 := []int{28,6,22,8,44,17}
	arr2 := []int{22,28,8,6}
	fmt.Println(relativeSortArray(arr1, arr2))

	arr3 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr4 := []int{2,1,4,3,9,6}
	fmt.Println(relativeSortArray(arr3, arr4))
}