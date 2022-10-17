package main

import (
	"fmt"
)

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
	fmt.Println(asteroidCollision([]int{8, -8}))
	fmt.Println(asteroidCollision([]int{10, 2, -5}))
	fmt.Println(asteroidCollision([]int{-2, -1, 1, 2}))
}

func asteroidCollision(asteroids []int) []int {
	s := NewStack()
	for _, a := range asteroids {
		for {
			if s.IsEmpty() {
				s.Push(a)
				break
			}

			la := s.Pop()
			if !isCollidable(la, a) {
				s.Push(la)
				s.Push(a)
				break
			}
			if abs(a) < abs(la) {
				s.Push(la)
				break
			}
			if abs(a) == abs(la) {
				break
			}
		}
	}
	return s.Value()
}

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: []int{},
	}
}

func (s *Stack) Value() []int {
	return s.data
}

func (s *Stack) Print() {
	fmt.Println(s.data)
}

func (s *Stack) Push(d int) {
	s.data = append(s.data, d)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	d := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return d
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func isCollidable(a, b int) bool {
	if a > 0 && b < 0 {
		return true
	}
	return false
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
