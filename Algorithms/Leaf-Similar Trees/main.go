package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	l := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	r := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root := TreeNode{
		Val:   3,
		Left:  &l,
		Right: &r,
	}
	fmt.Println(leafSimilar(&root, &root))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	var leave1, leave2 []int
	leave1 = getTreeLeaves(root1, leave1)
	leave2 = getTreeLeaves(root2, leave2)
	if len(leave1) != len(leave2) {
		return false
	}
	for i := 0; i < len(leave1); i++ {
		if leave1[i] != leave2[i] {
			return false
		}
	}
	return true
}

func getTreeLeaves(root *TreeNode, leaves []int) []int {
	if root == nil {
		return leaves
	}
	if root.Right == nil && root.Left == nil {
		leaves = append(leaves, root.Val)
	} else {
		leaves = getTreeLeaves(root.Left, leaves)
		leaves = getTreeLeaves(root.Right, leaves)
	}
	return leaves
}
