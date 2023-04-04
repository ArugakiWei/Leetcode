package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Self  *TreeNode
	Super *Node
}

type Queue struct {
	Nodes []*Node
}

func NewQueue() *Queue {
	return &Queue{
		Nodes: make([]*Node, 0, 0),
	}
}

func (q *Queue) push(node *Node) {
	q.Nodes = append(q.Nodes, node)
}

func (q *Queue) pop() *Node {
	p := q.Nodes[0]
	q.Nodes = q.Nodes[1:]
	return p
}

func (q *Queue) Len() int {
	return len(q.Nodes)
}

func sumEvenGrandparent(root *TreeNode) int {
	q := NewQueue()
	q.push(&Node{
		Self: root,
	})
	sum := 0
	for q.Len() > 0 {
		l := q.Len()
		for i := 0; i < l; i++ {
			n := q.pop()
			if n.Super != nil && n.Super.Super != nil && n.Super.Super.Self.Val%2 == 0 {
				sum += n.Self.Val
			}
			if n.Self.Left != nil {
				q.push(&Node{
					Self:  n.Self.Left,
					Super: n,
				})
			}
			if n.Self.Right != nil {
				q.push(&Node{
					Self:  n.Self.Right,
					Super: n,
				})
			}
		}
	}
	return sum
}
