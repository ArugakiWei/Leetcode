package main

import "fmt"

func main() {
	//fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4))
}

type MonotonicQueue struct {
	nums []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		nums: []int{},
	}
}

func (m *MonotonicQueue) push(i int) {
	for !m.isEmpty() && m.back() < i {
		m.popBack()
	}
	m.pushBack(i)
}

func (m *MonotonicQueue) max() int {
	return m.front()
}

func (m *MonotonicQueue) pop(n int) {
	if !m.isEmpty() && m.front() == n {
		m.popFront()
	}
}

func (m *MonotonicQueue) isEmpty() bool {
	return len(m.nums) == 0
}

func (m *MonotonicQueue) back() int {
	return m.nums[len(m.nums)-1]
}

func (m *MonotonicQueue) front() int {
	return m.nums[0]
}

func (m *MonotonicQueue) popBack() {
	m.nums = m.nums[:len(m.nums)-1]
}

func (m *MonotonicQueue) popFront() {
	m.nums = m.nums[1:]
}

func (m *MonotonicQueue) pushBack(i int) {
	m.nums = append(m.nums, i)
}

func maxSlidingWindow(nums []int, k int) []int {
	window, res := NewMonotonicQueue(), []int{}

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			window.push(nums[i])
		} else {
			window.push(nums[i])
			res = append(res, window.max())
			window.pop(nums[i-k+1])
		}
	}
	return res
}
