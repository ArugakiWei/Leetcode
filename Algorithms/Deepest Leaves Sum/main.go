package main

import "fmt"

func main() {
	a := &TreeNode{
		Val: 3,
	}
	b := &TreeNode{
		Val: 2,
	}
	r := &TreeNode{
		Val:   1,
		Left:  a,
		Right: b,
	}
	fmt.Println(deepestLeavesSum(r))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	Nodes []*TreeNode
}

func NewQueue() *Queue {
	return &Queue{
		Nodes: make([]*TreeNode, 0, 0),
	}
}

func (q *Queue) push(node *TreeNode) {
	q.Nodes = append(q.Nodes, node)
}

func (q *Queue) pop() *TreeNode {
	p := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return p
}

func (q *Queue) Len() int {
	return len(q.Nodes)
}

func deepestLeavesSum(root *TreeNode) int {
	q := NewQueue()
	q.push(root)
	sum := 0
	for q.Len() > 0 {
		sumLevel := 0
		l := q.Len()
		for i := 0; i < l; i++ {
			node := q.pop()
			if node.Left != nil {
				q.push(node.Left)
			}
			if node.Right != nil {
				q.push(node.Right)
			}
			sumLevel += node.Val
		}
		sum = sumLevel
	}
	return sum
}
