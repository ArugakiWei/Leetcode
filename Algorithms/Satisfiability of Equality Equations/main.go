package main

import "fmt"

type UF struct {
	Count   int   // 连通分量
	Parents []int // 树节点
	Size    []int // 节点重量
}

func NewUF(n int) *UF {
	u := &UF{
		Count:   n,
		Parents: make([]int, n),
		Size:    make([]int, n),
	}
	for i, _ := range u.Parents {
		u.Parents[i] = i
		u.Size[i] = 1
	}
	return u
}

func (u *UF) union(p, q int) {
	rootP := u.find(p)
	rootQ := u.find(q)
	if rootP == rootQ {
		return
	}

	if u.Size[q] > u.Size[p] {
		u.Parents[rootQ] = rootP
		u.Size[rootP] += u.Size[rootQ]
	} else {
		u.Parents[rootP] = rootQ
		u.Size[rootQ] += u.Size[rootP]
	}

	u.Count--
}

func (u *UF) find(x int) int {
	for u.Parents[x] != x {
		// 进行路径压缩
		u.Parents[x] = u.Parents[u.Parents[x]]
		x = u.Parents[x]
	}
	return x
}

func (u *UF) connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UF) count() int {
	return u.Count
}

func equationsPossible(equations []string) bool {
	u := NewUF(26)

	for _, v := range equations {
		if v[1] == '=' {
			u.union(int(v[0]-'a'), int(v[3]-'a'))
		}
	}

	for _, v := range equations {
		if v[1] == '!' {
			if u.connected(int(v[0]-'a'), int(v[3]-'a')) {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))
	fmt.Println(equationsPossible([]string{"b==a", "a==b"}))
}
