package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal1(root *TreeNode) []int {
	var res []int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		res = append(res, root.Val)
		traverse(root.Right)
	}
	traverse(root)
	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := newStack()
	p := root
	for !stack.isEmtpy() || p != nil {
		for p != nil {
			stack.push(p)
			p = p.Left
		}
		t := stack.pop()
		res = append(res, t.Val)
		if t.Right != nil {
			p = t.Right
		}
	}
	return res
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
