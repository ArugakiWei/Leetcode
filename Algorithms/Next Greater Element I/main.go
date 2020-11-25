package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}

//func nextGreaterElement(nums1 []int, nums2 []int) []int {
//	m := make(map[int]int)
//	for i, j := 0, 0; i < len(nums2)-1; i++ {
//		if nums2[i] < nums2[i+1] {
//			if j != 0 {
//				x := i - 1
//				for x >= 0 {
//					_, ok := m[nums2[x]]
//					if nums2[x] < nums2[i+1] && !ok {
//						m[nums2[x]] = nums2[i+1]
//						j--
//					}
//					x--
//				}
//			}
//			m[nums2[i]] = nums2[i+1]
//		} else {
//			j++
//		}
//	}
//
//	res := []int{}
//	for _, v := range nums1 {
//		if V, ok := m[v]; ok {
//			res = append(res, V)
//		} else {
//			res = append(res, -1)
//		}
//	}
//	return res
//}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	stack := NewStack()
	for _, v := range nums2 {
		for !stack.IsEmpty() && stack.top() < v {
			m[stack.pop()] = v
		}
		stack.push(v)
	}

	res := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		if V, ok := m[v]; ok {
			res = append(res, V)
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
