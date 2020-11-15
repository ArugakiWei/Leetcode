package main

import (
	"container/list"
	"fmt"
)

func main() {
	d := &TreeNode{
		Val: 5,
	}
	c := &TreeNode{
		Val: 4,
		Right: d,
	}
	b := &TreeNode{
		Val: 3,
		Right: c,
	}
	a := &TreeNode{
		Val: 2,
		Right: b,
	}
	root := &TreeNode{
		Val: 1,
		Right: a,
	}

	fmt.Println(minDepth(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := list.New()
	q.PushBack(root)
	depth := 1

	for q.Len() > 0 {
		n := q.Len()
		for i:=0; i<n; i++ {
			curElement := q.Front()
			q.Remove(curElement)

			cur := curElement.Value.(*TreeNode)
			if cur == nil {
				return depth
			}
			if cur.Left == nil && cur.Right == nil {
				return depth
			}

			if cur.Left != nil {
				q.PushBack(cur.Left)
			}
			if cur.Right != nil {
				q.PushBack(cur.Right)
			}
		}
		depth++
	}

	return depth
}
