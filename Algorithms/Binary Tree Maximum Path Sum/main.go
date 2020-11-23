package main

import (
	"fmt"
)

func main() {
	x := &TreeNode{
		Val: -1,
	}
	z := &TreeNode{
		Val: -2,
	}
	v := &TreeNode{
		Val:  1,
		Left: x,
	}
	i := &TreeNode{
		Val: 3,
	}
	a := &TreeNode{
		Val:   -2,
		Left:  v,
		Right: i,
	}
	b := &TreeNode{
		Val:  -3,
		Left: z,
	}
	c := &TreeNode{
		Val:   1,
		Left:  a,
		Right: b,
	}
	fmt.Println(maxPathSum(c))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	Ans int
}

func maxPathSum(root *TreeNode) int {
	ans := &Result{
		Ans: -1 << 31,
	}
	maxPathSum1(root, ans)
	return ans.Ans
}

func maxPathSum1(root *TreeNode, ans *Result) int {
	if root == nil {
		return 0
	}

	left := max(0, maxPathSum1(root.Left, ans))
	right := max(0, maxPathSum1(root.Right, ans))

	t := left + right + root.Val
	if ans.Ans < t {
		ans.Ans = t
	}

	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
