package main

import "fmt"

func main() {
	a := &TreeNode{
		Val: 0,
	}
	b := &TreeNode{
		Val: 1,
	}
	c := &TreeNode{
		Val:   3,
		Left:  a,
		Right: b,
	}
	d := &TreeNode{
		Val: 1,
	}
	e := &TreeNode{
		Val:   2,
		Left:  d,
		Right: c,
	}
	fmt.Println(evaluateTree(e))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if isLeaf(root) {
		return val(root.Val)
	}
	if isLeaf(root.Left) && isLeaf(root.Right) {
		return calculate(root.Val, val(root.Left.Val), val(root.Right.Val))
	}
	return calculate(root.Val, evaluateTree(root.Left), evaluateTree(root.Right))
}

func val(v int) bool {
	return v == 1
}

func calculate(rootv int, lv, rv bool) bool {
	if rootv == 2 {
		return lv || rv
	}
	return lv && rv
}

func isLeaf(n *TreeNode) bool {
	return n.Left == nil && n.Right == nil
}
