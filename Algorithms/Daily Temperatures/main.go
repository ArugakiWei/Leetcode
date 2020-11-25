package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

type Stack struct {
	nums []int
}

func NewStack() *Stack {
	return &Stack{
		nums: []int{},
	}
}

func (s *Stack) pop() int {
	if len(s.nums) == 0 {
		return -1
	}

	top := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return top
}

func (s *Stack) push(i int) {
	s.nums = append(s.nums, i)
}

func (s *Stack) IsEmpty() bool {
	return len(s.nums) == 0
}

func (s *Stack) top() int {
	return s.nums[len(s.nums)-1]
}

func dailyTemperatures(T []int) []int {
	res, stack := make([]int, len(T)), NewStack()
	for i := len(T) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && T[stack.top()] <= T[i] {
			stack.pop()
		}
		if !stack.IsEmpty() {
			res[i] = stack.top() - i
		} else {
			res[i] = 0
		}
		stack.push(i)
	}

	return res
}
