package main

import (
	"fmt"
)

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}

func numIslands1(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n {
			return
		} // 递归终止条件：越界（无效的网格），则返回
		if grid[r][c] != '1' {
			return
		} // 非未访问过的陆地（水or已访问过的陆地），则返回
		grid[r][c] = '2' // 标记已访问过的陆地：防止无限循环无法退出
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' { // 遇到一个陆地，结果+1，并则将其周围四个相连的陆地变为非陆地
				ans++
				dfs(r, c)
			}
		}
	}
	return
}

type UF struct {
	count   int
	parents map[string]string
}

func newUF(grid [][]byte) *UF {
	n, m, count := len(grid), len(grid[0]), 0
	parents := make(map[string]string, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				count++
				parents[val(i, j)] = val(i, j)
			}
		}
	}
	return &UF{
		count:   count,
		parents: parents,
	}
}

func val(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

func (u *UF) find(point string) string {
	if u.parents[point] != point {
		u.parents[point] = u.find(u.parents[point])
	}
	return u.parents[point]
}

func (u *UF) union(p, q string) {
	rootP := u.find(p)
	rootQ := u.find(q)
	if rootQ == rootP {
		return
	}

	u.parents[rootQ] = rootP
	u.count--
	return
}

func (u *UF) isConnected(p, q string) bool {
	return u.find(p) == u.find(q)
}

func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	uf := newUF(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				if i-1 >= 0 && grid[i-1][j] == '1' {
					uf.union(val(i-1, j), val(i, j))
					continue
				}
				if i+1 < n && grid[i+1][j] == '1' {
					uf.union(val(i+1, j), val(i, j))
					continue
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					uf.union(val(i, j-1), val(i, j))
					continue
				}
				if j+1 < m && grid[i][j+1] == '1' {
					uf.union(val(i, j+1), val(i, j))
					continue
				}
			}
		}
	}
	return uf.count
}
