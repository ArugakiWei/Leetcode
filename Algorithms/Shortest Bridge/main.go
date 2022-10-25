package main

import "fmt"

func main() {
	fmt.Println(shortestBridge([][]int{
		{0, 1},
		{1, 0},
	}))
	fmt.Println(shortestBridge([][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 0, 1},
	}))
	fmt.Println(shortestBridge([][]int{
		{1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 1, 1, 1},
	}))
	fmt.Println(shortestBridge([][]int{
		{0, 0, 1, 0, 1},
		{0, 1, 1, 0, 1},
		{0, 1, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
	fmt.Println(shortestBridge([][]int{
		{0, 1, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0},
	}))
}

type Point struct {
	X int
	Y int
}

func distance(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y) - 1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func isConnect(a *[]Point, b Point) bool {
	if len(*a) == 0 {
		return true
	}
	for _, p := range *a {
		if p.X == b.X && p.Y == b.Y {
			return true
		}
		if abs(p.X-b.X) == 1 && p.Y == b.Y {
			return true
		}
		if abs(p.Y-b.Y) == 1 && p.X == b.X {
			return true
		}
	}
	return false
}

func joinIsland(grid [][]int, island *[]Point, p Point, maxX, maxY int) {
	if p.X > maxX || p.Y > maxY || p.X < 0 || p.Y < 0 || grid[p.Y][p.X] != 1 {
		return
	}
	for _, ip := range *island {
		if ip.X == p.X && ip.Y == p.Y {
			return
		}
	}
	if isConnect(island, p) {
		*island = append(*island, p)
	}
	joinIsland(grid, island, Point{X: p.X - 1, Y: p.Y}, maxX, maxY)
	joinIsland(grid, island, Point{X: p.X + 1, Y: p.Y}, maxX, maxY)
	joinIsland(grid, island, Point{X: p.X, Y: p.Y - 1}, maxX, maxY)
	joinIsland(grid, island, Point{X: p.X, Y: p.Y + 1}, maxX, maxY)
}

func shortestBridge(grid [][]int) int {
	islandA := make([]Point, 0, 0)
	islandB := make([]Point, 0, 0)
	for i, v := range grid {
		for j, vv := range v {
			p := Point{
				X: j,
				Y: i,
			}
			// 构造a岛
			if vv == 1 && len(islandA) == 0 {
				joinIsland(grid, &islandA, p, len(v)-1, len(grid)-1)
			}
			if vv == 1 && !isConnect(&islandA, p) {
				islandB = append(islandB, p)
			}
		}
	}
	minDistance := distance(islandA[0], islandB[0])
	for _, ap := range islandA {
		for _, bp := range islandB {
			if distance(ap, bp) < minDistance {
				minDistance = distance(ap, bp)
			}
		}
	}
	return minDistance
}
