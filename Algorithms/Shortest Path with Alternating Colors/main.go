package main

import "fmt"

func main() {
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}, {1, 2}}, [][]int{}))
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{2, 1}}))
	fmt.Println(shortestAlternatingPaths(3, [][]int{{1, 0}}, [][]int{{2, 1}}))
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}}, [][]int{{1, 2}}))
	fmt.Println(shortestAlternatingPaths(3, [][]int{{0, 1}, {0, 2}}, [][]int{{1, 0}}))
	fmt.Println(shortestAlternatingPaths(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 2}, {2, 3}, {3, 1}}))
	fmt.Println(shortestAlternatingPaths(5, [][]int{{3, 2}, {4, 1}, {1, 4}, {2, 4}}, [][]int{{2, 3}, {0, 4}, {4, 3}, {4, 4}, {4, 0}, {1, 0}}))
	fmt.Println(shortestAlternatingPaths(5, [][]int{{2, 2}, {0, 1}, {0, 3}, {0, 0}, {0, 4}, {2, 1}, {2, 0}, {1, 4}, {3, 4}}, [][]int{{1, 3}, {0, 0}, {0, 3}, {4, 2}, {1, 0}}))
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	redGraph, blueGraph := make(map[int]map[int]struct{}), make(map[int]map[int]struct{})
	for _, re := range redEdges {
		if _, ok := redGraph[re[0]]; !ok {
			redGraph[re[0]] = make(map[int]struct{})
		}
		redGraph[re[0]][re[1]] = struct{}{}
	}
	for _, re := range blueEdges {
		if _, ok := blueGraph[re[0]]; !ok {
			blueGraph[re[0]] = make(map[int]struct{})
		}
		blueGraph[re[0]][re[1]] = struct{}{}
	}
	ans := make([]int, n)
	for i := 1; i < n; i++ {
		ans[i] = search(redGraph, blueGraph, i)
	}
	return ans
}

func search(redGraph, blueGraph map[int]map[int]struct{}, target int) int {
	m := make(map[string]struct{})

	q := newQueue()
	q.Push(&Point{
		Value: 0,
		IsRed: true,
	})
	q.Push(&Point{
		Value: 0,
		IsRed: false,
	})
	for !q.IsEmpty() {
		point := q.Pop()
		m[point.String()] = struct{}{}

		graph := redGraph
		if !point.IsRed {
			graph = blueGraph
		}

		for k, _ := range graph[point.Value] {
			if k == target {
				return point.Count + 1
			}
			kp := &Point{
				Value: k,
				IsRed: !point.IsRed,
				Count: point.Count + 1,
			}
			if _, ok := m[kp.String()]; !ok {
				q.Push(kp)
			}
		}
	}
	return -1
}

type Point struct {
	Value int
	IsRed bool
	Count int
}

func (p *Point) String() string {
	return fmt.Sprintf("%d-%v", p.Value, p.IsRed)
}

func newQueue() *Queue {
	return &Queue{
		Points: make([]*Point, 0, 0),
	}
}

type Queue struct {
	Points []*Point
}

func (q *Queue) Push(p *Point) {
	q.Points = append(q.Points, p)
}

func (q *Queue) Pop() *Point {
	p := q.Points[0]
	q.Points = q.Points[1:]
	return p
}

func (q *Queue) IsEmpty() bool {
	return len(q.Points) == 0
}
