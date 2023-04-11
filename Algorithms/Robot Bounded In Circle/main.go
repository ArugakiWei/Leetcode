package main

import "fmt"

func main() {
	fmt.Println(isRobotBounded("GL"))
}

func isRobotBounded(instructions string) bool {
	ins := instructions + instructions + instructions + instructions
	x, y := exec(ins, 0, 0)
	return x == 0 && y == 0
}

func exec(instructions string, x, y int) (int, int) {
	s := 'n'
	for _, v := range instructions {
		if v == 'G' {
			switch s {
			case 'n':
				y += 1
			case 'e':
				x += 1
			case 'w':
				x -= 1
			case 's':
				y -= 1
			}
		} else {
			s = direction(s, v)
		}
	}
	return x, y
}

func direction(cur, dst int32) int32 {
	switch cur {
	case 'n':
		if dst == 'L' {
			return 'w'
		}
		return 'e'
	case 'e':
		if dst == 'L' {
			return 'n'
		}
		return 's'
	case 'w':
		if dst == 'L' {
			return 's'
		}
		return 'n'
	case 's':
		if dst == 'L' {
			return 'e'
		}
		return 'w'
	}
	return cur
}
