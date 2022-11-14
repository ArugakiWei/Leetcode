package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestPathAllKeys([]string{"@.a.#", "###.#", "b.A.B"}))
	fmt.Println(shortestPathAllKeys([]string{"@..aA", "..B#.", "....b"}))
	fmt.Println(shortestPathAllKeys([]string{"@Aa"}))
	fmt.Println(shortestPathAllKeys([]string{"@...a", ".###A", "b.BCc"}))
	fmt.Println(shortestPathAllKeys([]string{"@abcdeABCDEFf"}))
	fmt.Println(shortestPathAllKeys([]string{".#.#..#.b...............#.#..#", ".#..##.........#......d.......", "..#...e.#.##....##.....#.....#", "..#..#.#.#.##..........#.....#", "...#...##....#.....#..........", "#........###....#..#.........f", "...............#......#...#...", "..........##.#...#.E..#......#", ".#...##...#.##.D....##..#.....", ".......#...........#....#..##.", "...#..........##.....#.......#", ".F#....#......#...............", "..##.#.#.....#..##...#.#.....#", ".............##..##..#.#......", "#..@..#.#.......#..........#..", ".........##..................#", ".#.......##...##..#.......#...", ".......#.#...A.a......#.##.#..", ".......#......##..#.###.#.....", ".##.#....##...#.#.....#.#.....", ".#.....#.c..#.....#......#..##", "##.....##........B.#.......#.#", ".....#...#....#..##...........", "#.#.##.#....#.#...............", ".#.#..#.####............#.....", "#.#..........###.#........#...", "..#..#.........#.......#..#.##", "..#..#C#...............#......", ".........#.##.##......#.#.....", "..#........##.#..##.#.....#.#."}))
}

// @...a
// .###A
// b.BCc

func shortestPathAllKeys(grid []string) int {
	startX, startY, keyCount := 0, 0, 0
	for i, v := range grid {
		for j, vv := range v {
			if vv == '@' {
				startX, startY = i, j
			}
			if vv >= 97 && vv <= 122 {
				keyCount++
			}
		}
	}
	canGo := func(x, y int, p *Point) bool {
		c := grid[x][y]
		if c == '#' {
			return false
		}
		if c >= 65 && c <= 90 {
			if _, ok := p.Keys[c+32]; !ok {
				return false
			}
		}
		return true
	}
	getNewKeys := func(oldKeys map[byte]struct{}, c byte) map[byte]struct{} {
		newKeys := make(map[byte]struct{})
		for k, v := range oldKeys {
			newKeys[k] = v
		}
		// 拿到钥匙
		if c >= 97 && c <= 122 {
			newKeys[c] = struct{}{}
		}
		return newKeys
	}

	q := NewQueue(&Point{
		X:    startX,
		Y:    startY,
		Keys: make(map[byte]struct{}),
	})
	step := 0
	for {
		// 没有拿到全部钥匙先无路可走
		if q.IsEmpty() {
			return -1
		}

		size := q.Len()
		for i := 0; i < size; i++ {
			p := q.Dequeue()
			// 有某一路径上拿到所有的钥匙
			if p.KeyCount() == keyCount {
				return step
			}
			// 向上
			if p.X-1 >= 0 && canGo(p.X-1, p.Y, p) {
				kc := getNewKeys(p.Keys, grid[p.X-1][p.Y])
				if len(kc) == keyCount {
					return step + 1
				}
				q.Enqueue(&Point{
					X:    p.X - 1,
					Y:    p.Y,
					Keys: kc,
				})
			}
			// 向下
			if p.X+1 < len(grid) && canGo(p.X+1, p.Y, p) {
				kc := getNewKeys(p.Keys, grid[p.X+1][p.Y])
				if len(kc) == keyCount {
					return step + 1
				}
				q.Enqueue(&Point{
					X:    p.X + 1,
					Y:    p.Y,
					Keys: getNewKeys(p.Keys, grid[p.X+1][p.Y]),
				})
			}
			// 向左
			if p.Y-1 >= 0 && canGo(p.X, p.Y-1, p) {
				kc := getNewKeys(p.Keys, grid[p.X][p.Y-1])
				if len(kc) == keyCount {
					return step + 1
				}
				q.Enqueue(&Point{
					X:    p.X,
					Y:    p.Y - 1,
					Keys: getNewKeys(p.Keys, grid[p.X][p.Y-1]),
				})
			}
			// 向右
			if p.Y+1 < len(grid[p.X]) && canGo(p.X, p.Y+1, p) {
				kc := getNewKeys(p.Keys, grid[p.X][p.Y+1])
				if len(kc) == keyCount {
					return step + 1
				}
				q.Enqueue(&Point{
					X:    p.X,
					Y:    p.Y + 1,
					Keys: getNewKeys(p.Keys, grid[p.X][p.Y+1]),
				})
			}
		}
		step++
	}
}

type Point struct {
	X    int
	Y    int
	Keys map[byte]struct{}
}

func (p *Point) KeyCount() int {
	return len(p.Keys)
}

func (p *Point) String() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("(%d, %d, %v)", p.X, p.Y, p.Keys)
}

func NewQueue(p *Point) *Queue {
	return &Queue{
		data: []*Point{p},
		mem:  make(map[string]*Point),
	}
}

type Queue struct {
	data []*Point
	mem  map[string]*Point
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Enqueue(p *Point) {
	_, ok := q.mem[p.String()]
	if !ok {
		q.mem[p.String()] = p
		q.data = append(q.data, p)
		return
	}
}

func (q *Queue) Dequeue() *Point {
	d := q.data[0]
	q.data = q.data[1:]
	return d
}

func (q *Queue) Range() []*Point {
	return q.data
}

func (q *Queue) Len() int {
	return len(q.data)
}
