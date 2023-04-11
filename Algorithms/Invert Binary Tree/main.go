package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Right)
	invertTree(root.Left)
	root.Right, root.Left = root.Left, root.Right
	return root
}
