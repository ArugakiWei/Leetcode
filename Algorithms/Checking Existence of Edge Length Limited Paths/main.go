package main

import "fmt"

func main() {
	// fmt.Println(distanceLimitedPathsExist(3, [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}, [][]int{{0, 1, 2}, {0, 2, 5}}))
	// fmt.Println(distanceLimitedPathsExist(5, [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}}, [][]int{{0, 4, 14}, {1, 4, 13}}))
	fmt.Println(distanceLimitedPathsExist(13, [][]int{{9, 1, 53}, {3, 2, 66}, {12, 5, 99}, {9, 7, 26}, {1, 4, 78}, {11, 1, 62}, {3, 10, 50}, {12, 1, 71}, {12, 6, 63}, {1, 10, 63}, {9, 10, 88}, {9, 11, 59}, {1, 4, 37}, {4, 2, 63}, {0, 2, 26}, {6, 12, 98}, {9, 11, 99}, {4, 5, 40}, {2, 8, 25}, {4, 2, 35}, {8, 10, 9}, {11, 9, 25}, {10, 11, 11}, {7, 6, 89}, {2, 4, 99}, {10, 4, 63}}, [][]int{{3, 7, 76}}))
}

type Node struct {
	Edge   map[int]int
	Number int
	Nodes  map[int]*Node
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	m := make(map[int]*Node)
	for _, edge := range edgeList {
		p0, p1, w := edge[0], edge[1], edge[2]
		node0, ok0 := m[p0]
		node1, ok1 := m[p1]
		if !ok0 && !ok1 {
			node0, node1 = NewNode(p0), NewNode(p1)
			node0.Nodes[p1], node1.Nodes[p0] = node1, node0
			node0.Edge[p1], node1.Edge[p0] = w, w
			m[p0], m[p1] = node0, node1
		}
		if ok0 && !ok1 {
			node1 = NewNode(p1)
			node0.Nodes[p1], node1.Nodes[p0] = node1, node0
			node0.Edge[p1], node1.Edge[p0] = w, w
			m[p1] = node1
		}
		if !ok0 && ok1 {
			node0 = NewNode(p0)
			node0.Nodes[p1], node1.Nodes[p0] = node1, node0
			node0.Edge[p1], node1.Edge[p0] = w, w
			m[p0] = node0
		}
		if ok0 && ok1 {
			setPathWeight(m[p0], p1, w)
			setPathWeight(m[p1], p0, w)
			node0.Nodes[p1], node1.Nodes[p0] = node1, node0
		}
	}

	var ans []bool
	for _, query := range queries {
		ans = append(ans, checkQuery(query, m))
	}
	return ans
}

func checkQuery(query []int, m map[int]*Node) bool {
	start := query[0]
	end := query[1]
	limit := query[2]
	mem := make(map[int]struct{})
	q := NewQueue()
	q.Push(m[start])
	cc := 0
	for q.Len() > 0 {
		l := q.Len()
		cc++
		for i := 0; i < l; i++ {
			x := q.Pop()
			mem[x.Number] = struct{}{}
			if w, ok := x.Edge[end]; ok && w < limit {
				return true
			}
			for number, weight := range x.Edge {
				_, ok1 := mem[number]
				if weight < limit && !ok1 {
					q.Push(x.Nodes[number])
				}
			}
		}
	}
	return false
}

func NewQueue() *Queue {
	return &Queue{
		Data: make([]*Node, 0, 0),
	}
}

type Queue struct {
	Data []*Node
}

func (q *Queue) Len() int {
	return len(q.Data)
}

func (q *Queue) Push(n *Node) {
	if n == nil {
		return
	}
	q.Data = append(q.Data, n)
}

func (q *Queue) Pop() *Node {
	if q.Len() == 0 {
		return nil
	}
	p := q.Data[0]
	q.Data = q.Data[1:]
	return p
}

func setPathWeight(n *Node, p int, weight int) {
	if v, ok := n.Edge[p]; ok {
		if weight < v {
			n.Edge[p] = weight
		}
	} else {
		n.Edge[p] = weight
	}
}

func NewNode(n int) *Node {
	return &Node{
		Edge:   make(map[int]int),
		Number: n,
		Nodes:  make(map[int]*Node),
	}
}
