package main

import "fmt"

func main() {
	fmt.Println(allPathsSourceTarget([][]int{
		{1, 2},
		{3},
		{3},
		{},
	}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	ans, path := make([][]int, 0, 0), make([]int, 0, len(graph))
	path = append(path, 0)
	var backtrack func(cur int)
	backtrack = func(cur int) {
		if cur == n-1 {
			t := make([]int, len(path))
			copy(t, path)
			ans = append(ans, t)
			return
		}
		for i := 0; i < len(graph[cur]); i++ {
			path = append(path, graph[cur][i])
			backtrack(graph[cur][i])
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ans
}
