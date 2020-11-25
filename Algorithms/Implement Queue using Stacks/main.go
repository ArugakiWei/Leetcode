package main

type MyQueue struct {
	input  *Stack
	output *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		input:  NewStack(),
		output: NewStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.input.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.output.isEmpty() && this.input.isEmpty() {
		return -1
	}
	if this.output.isEmpty() && !this.input.isEmpty() {
		for !this.input.isEmpty() {
			x := this.input.pop()
			this.output.push(x)
		}
	}

	return this.output.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.output.isEmpty() && this.input.isEmpty() {
		return -1
	}
	if this.output.isEmpty() && !this.input.isEmpty() {
		for !this.input.isEmpty() {
			x := this.input.pop()
			this.output.push(x)
		}
	}

	return this.output.top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.input.isEmpty() && this.output.isEmpty()
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

func (s *Stack) isEmpty() bool {
	return len(s.nums) == 0
}

func (s *Stack) top() int {
	return s.nums[len(s.nums)-1]
}
