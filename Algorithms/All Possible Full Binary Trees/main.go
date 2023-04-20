package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return []*TreeNode{}
	}
	d := [20][]*TreeNode{
		1: {
			&TreeNode{},
		},
	}
	for n := 3; n <= N; n += 2 {
		for i := 1; i < n; i += 2 {
			for _, left := range d[i] {
				for _, right := range d[n-i-1] {
					d[n] = append(d[n], &TreeNode{Left: left, Right: right})
				}
			}
		}
	}
	return d[N]
}
