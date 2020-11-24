package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 两个都为空
	if p == nil && q == nil {
		return true
	}
	// 一个为空, 另一个不为空
	if p == nil || q == nil {
		return false
	}
	// 值不相等
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
