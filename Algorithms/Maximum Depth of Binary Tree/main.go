package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	ans, depth := 0, 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		depth++
		if root.Right == nil && root.Left == nil {
			ans = max(ans, depth)
		}
		dfs(root.Right)
		dfs(root.Left)
		depth--
	}
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
