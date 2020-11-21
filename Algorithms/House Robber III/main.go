package main

import "fmt"

func main() {
	a := &TreeNode{
		Val: 3,
	}
	b := &TreeNode{
		Val: 1,
	}
	c := &TreeNode{
		Val:   2,
		Right: a,
	}
	d := &TreeNode{
		Val:   3,
		Right: b,
	}
	e := &TreeNode{
		Val:   3,
		Left:  c,
		Right: d,
	}
	fmt.Println(rob(e))

	a1 := &TreeNode{
		Val: 1,
	}
	a2 := &TreeNode{
		Val: 3,
	}
	b1 := &TreeNode{
		Val: 1,
	}
	c1 := &TreeNode{
		Val:   4,
		Left:  a1,
		Right: a2,
	}
	d1 := &TreeNode{
		Val:   5,
		Right: b1,
	}
	e1 := &TreeNode{
		Val:   3,
		Left:  c1,
		Right: d1,
	}
	fmt.Println(rob(e1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	return robMemo(root, make(map[interface{}]int))
}

func robMemo(root *TreeNode, memo map[interface{}]int) int {
	if root == nil {
		return 0
	}
	if v, ok := memo[root]; ok {
		return v
	}

	doRob, dontRob, res := root.Val, robMemo(root.Left, memo)+robMemo(root.Right, memo), 0
	if root.Right != nil && root.Left != nil {
		doRob = root.Val + robMemo(root.Left.Left, memo) + robMemo(root.Left.Right, memo) + robMemo(root.Right.Right, memo) + robMemo(root.Right.Left, memo)
	}
	if root.Right == nil && root.Left != nil {
		doRob = root.Val + robMemo(root.Left.Left, memo) + robMemo(root.Left.Right, memo)
	}
	if root.Right != nil && root.Left == nil {
		doRob = root.Val + robMemo(root.Right.Right, memo) + robMemo(root.Right.Left, memo)
	}

	res = max(doRob, dontRob)
	memo[root] = res

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
