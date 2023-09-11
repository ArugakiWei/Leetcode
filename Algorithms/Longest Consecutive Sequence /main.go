package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

type UF struct {
	parents map[int]int
}

func newUF(nums []int) *UF {
	parents := make(map[int]int)
	for _, n := range nums {
		parents[n] = n
	}
	return &UF{
		parents: parents,
	}
}

func (u *UF) union(p, q int) {
	rootP, rootQ := u.find(p), u.find(q)
	if rootQ == rootP {
		return
	}
	u.parents[rootQ] = rootP
}

func (u *UF) find(p int) int {
	if u.parents[p] != p {
		u.parents[p] = u.find(u.parents[p])
	}
	return u.parents[p]
}

func longestConsecutive(nums []int) int {
	uf := newUF(nums)
	for _, n := range nums {
		if _, ok := uf.parents[n+1]; ok {
			uf.union(n+1, n)
			continue
		}
		if _, ok := uf.parents[n-1]; ok {
			uf.union(n-1, n)
			continue
		}
	}
	count := make(map[int]int)
	for k, _ := range uf.parents {
		root := uf.find(k)
		count[root]++
	}
	ans := 0
	for _, c := range count {
		if c >= ans {
			ans = c
		}
	}
	return ans
}

func longestConsecutive1(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
