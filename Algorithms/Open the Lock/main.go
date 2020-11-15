package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(openLock([]string{"0201","0101","0102","1212","2002"}, "0202"))
	fmt.Println(openLock([]string{"8888"}, "0009"))
	fmt.Println(openLock([]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, "8888"))
	fmt.Println(openLock([]string{"0000"}, "8888"))
}

func openLock(deadends []string, target string) int {
	visited := make(map[string]struct{})
	for _, deadend := range deadends {
		visited[deadend] = struct{}{}
	}
	if _, ok := visited[target]; ok {
		return -1
	}
	if _, ok := visited["0000"]; ok {
		return -1
	}

	q := list.New()
	q.PushBack("0000")

	step := 0

	for q.Len() > 0 {
		n := q.Len()
		for i:=0; i<n; i++ {
			curElement := q.Front()
			q.Remove(curElement)

			cur := curElement.Value.(string)
			if cur == target {
				return step
			}

			for i:=0; i<4; i++ {
				u := up(cur, i)
				if _, ok := visited[u]; !ok {
					visited[u] = struct{}{}
					q.PushBack(u)
				}

				d := down(cur, i)
				if _, ok := visited[d]; !ok {
					visited[d] = struct{}{}
					q.PushBack(d)
				}
			}
		}
		step++
	}
	return -1
}

func up(s string, p int) string {
	n := []byte(s)
	if n[p] == '9' {
		n[p] = '0'
	} else {
		n[p] += 1
	}
	return string(n)
}

func down(s string, p int) string {
	n := []byte(s)
	if n[p] == '0' {
		n[p] = '9'
	} else {
		n[p] -= 1
	}
	return string(n)
}