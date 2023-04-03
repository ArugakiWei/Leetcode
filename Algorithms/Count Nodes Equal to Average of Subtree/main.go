package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum, count := childrenSum(root)
	if int(math.Floor(float64(sum)/float64(count))) == root.Val {
		return 1 + averageOfSubtree(root.Left) + averageOfSubtree(root.Right)
	}
	return averageOfSubtree(root.Left) + averageOfSubtree(root.Right)
}

func childrenSum(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lSum, lCount := childrenSum(root.Left)
	rSum, rCount := childrenSum(root.Right)
	return root.Val + lSum + rSum, 1 + lCount + rCount
}
