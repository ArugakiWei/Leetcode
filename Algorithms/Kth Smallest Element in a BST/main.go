package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var ans int
	n := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		n++
		if n == k {
			ans = root.Val
			return
		}
		traverse(root.Right)
	}
	traverse(root)
	return ans
}
