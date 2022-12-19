package main

import "fmt"

func main() {
	// fmt.Println(validPath(3, [][]int{{1, 0}, {1, 2}, {2, 0}}, 0, 2))
	// fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
	fmt.Println(validPath(10, [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}, 7, 5))
}

func validPath1(n int, edges [][]int, source int, destination int) bool {
	if n == 1 && source == 0 && destination == 0 {
		return true
	}

	var graph []map[int]struct{}
	for _, v := range edges {
		g1 := inWhichGraph(graph, v[0])
		g2 := inWhichGraph(graph, v[1])
		// 2个点都不在任一集合中, 则2个点加入一个新集合
		if g1 == -1 && g2 == -1 {
			m := map[int]struct{}{
				v[0]: {},
				v[1]: {},
			}
			graph = append(graph, m)
			continue
		}
		// 一个点不在任意集合中, 另一个点已经在一个集合中
		// 加入另一个点的集合
		if g1 == -1 && g2 != -1 {
			graph[g2][v[0]] = struct{}{}
			continue
		}
		if g1 != -1 && g2 == -1 {
			graph[g1][v[1]] = struct{}{}
			continue
		}
		// 两个点都在了已经存在了相同集合中
		if g1 == g2 {
			continue
		}
		// 合并g1 g2
		if g1 != g2 {
			graph = mergeGraph(graph, g1, g2)
		}
	}
	for _, g := range graph {
		_, ok1 := g[source]
		_, ok2 := g[destination]
		if ok1 && ok2 {
			return true
		}
	}
	return false
}

func inWhichGraph(graph []map[int]struct{}, p int) int {
	for i, g := range graph {
		if _, ok := g[p]; ok {
			return i
		}
	}
	return -1
}

func mergeGraph(graph []map[int]struct{}, g1, g2 int) []map[int]struct{} {
	var newGraph []map[int]struct{}
	ng := make(map[int]struct{})
	for i, g := range graph {
		if i == g1 || i == g2 {
			for p := range graph[i] {
				ng[p] = struct{}{}
			}
			continue
		}
		newGraph = append(newGraph, g)
	}
	newGraph = append(newGraph, ng)
	return newGraph
}

func validPath2(n int, edges [][]int, start int, end int) bool {
	if start == end {
		return true
	}
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parents[x] == x {
			return x
		}
		parents[x] = find(parents[x])
		return parents[x]
	}
	for _, e := range edges {
		u := find(e[0])
		v := find(e[1])
		if u != v {
			parents[u] = v
		}
	}
	return find(start) == find(end)
}
