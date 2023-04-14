package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	Data []*TreeNode
}

func newStack() *Stack {
	return &Stack{
		Data: make([]*TreeNode, 0, 0),
	}
}

func (s *Stack) isEmtpy() bool {
	return len(s.Data) == 0
}

func (s *Stack) push(n *TreeNode) {
	s.Data = append(s.Data, n)
}
func (s *Stack) pop() *TreeNode {
	i := len(s.Data) - 1
	n := s.Data[i]
	s.Data = s.Data[:i]
	return n
}

func postorderTraversal1(root *TreeNode) []int {
	var res []int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		traverse(root.Right)
		res = append(res, root.Val)
	}
	traverse(root)
	return res
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack1 := newStack()
	stack2 := newStack()
	stack1.push(root)
	for !stack1.isEmtpy() {
		t := stack1.pop()
		stack2.push(t)
		if t.Left != nil {
			stack1.push(t.Left)
		}
		if t.Right != nil {
			stack1.push(t.Right)
		}
	}
	for !stack2.isEmtpy() {
		res = append(res, stack2.pop().Val)
	}
	return res
}
