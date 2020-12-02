package main

import "fmt"

func main() {
	fmt.Println(isValid("[](){}"))
	fmt.Println(isValid(")}"))
}

func isValid(s string) bool {
	stack := NewStack()
	for _, b := range []byte(s) {
		if b == '{' || b == '(' || b == '[' {
			stack.push(b)
		} else if !stack.isEmpty() {
			n := stack.pop()
			switch b {
			case '}':
				if n != '{' {
					return false
				}
			case ')':
				if n != '(' {
					return false
				}
			case ']':
				if n != '[' {
					return false
				}
			}
		} else {
			return false
		}
	}

	return stack.isEmpty()
}

type Stack struct {
	num []byte
}

func NewStack() *Stack {
	return &Stack{
		num: make([]byte, 0, 0),
	}
}

func (s *Stack) push(n byte) {
	s.num = append(s.num, n)
}

func (s *Stack) pop() byte {
	n := s.num[len(s.num)-1]
	s.num = s.num[:len(s.num)-1]
	return n
}

func (s *Stack) isEmpty() bool {
	return len(s.num) == 0
}
