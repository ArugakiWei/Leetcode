package main

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	R    *rand.Rand
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
		R:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) GetRandom() int {
	i, res := 0, 0
	t := this.head
	for t != nil {
		i++
		if this.R.Intn(i) == 0 {
			res = t.Val
		}
		t = t.Next
	}

	return res
}
