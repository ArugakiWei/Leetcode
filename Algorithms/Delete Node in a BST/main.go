package main

import "fmt"

func main() {
	a := &TreeNode{
		Val: 2,
	}
	b := &TreeNode{
		Val: 4,
	}
	c := &TreeNode{
		Val: 7,
	}
	d := &TreeNode{
		Val:   3,
		Left:  a,
		Right: b,
	}
	e := &TreeNode{
		Val:   6,
		Right: c,
	}
	f := &TreeNode{
		Val:   5,
		Left:  d,
		Right: e,
	}
	fmt.Println(deleteNode(f, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		// 没有子节点
		if root.Right == nil && root.Left == nil {
			return nil
		}
		// 只有一个子节点
		if root.Right == nil && root.Left != nil {
			return root.Left
		}
		if root.Right != nil && root.Left == nil {
			return root.Right
		}
		// 有2个子节点
		if root.Right != nil && root.Left != nil {
			// 找右子树中的最小节点, 交换其值, 然后删除
			minNode := getMinFromRight(root.Right)
			root.Val = minNode.Val
			root.Right = deleteNode(root.Right, minNode.Val)
		}
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMinFromRight(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}
