package main

import "fmt"

func main() {
	//fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	fmt.Println(nextGreaterElements([]int{100, 1, 11, 1, 120, 111, 123, 1, -1, -100}))
}

func nextGreaterElements(nums []int) []int {
	n := len(nums)

	m := make(map[int]int)
	stack := NewStack()
	for i := 0; i < 2*n-1; i++ {
		for !stack.IsEmpty() && nums[stack.top()] < nums[i%n] {
			m[stack.pop()] = i % n
		}
		stack.push(i % n)
	}

	res := make([]int, 0, len(nums))
	for i, _ := range nums {
		if ii, ok := m[i]; ok {
			res = append(res, nums[ii%len(nums)])
		} else {
			res = append(res, -1)
		}
	}
	return res
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
