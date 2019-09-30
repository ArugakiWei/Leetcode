package main

import "fmt"

func main() {
	fmt.Println()
}

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	if root != nil {
		if root.Val >= L && root.Val <= R {
			sum = sum + root.Val
		}
		if root.Left != nil {
			sum = sum + rangeSumBST(root.Left, L, R)
		}
		if root.Right != nil {
			sum = sum + rangeSumBST(root.Right, L, R)
		}
	}
	return sum
}