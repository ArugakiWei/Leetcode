package main

import "fmt"

func main() {
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	inorderIndex := 0
	for i, v := range inorder {
		if v == root.Val {
			inorderIndex = i
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:inorderIndex])+1], inorder[:inorderIndex])
	root.Right = buildTree(preorder[len(inorder[:inorderIndex])+1:], inorder[inorderIndex+1:])
	return root
}
