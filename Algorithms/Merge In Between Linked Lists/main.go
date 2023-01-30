package main

import "fmt"

func main() {
	a := &ListNode{Val: 0}
	b := &ListNode{Val: 1}
	c := &ListNode{Val: 2}
	d := &ListNode{Val: 3}
	e := &ListNode{Val: 4}
	f := &ListNode{Val: 5}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f

	g := &ListNode{Val: 10001}
	h := &ListNode{Val: 10002}
	k := &ListNode{Val: 10003}
	g.Next = h
	h.Next = k
	p(mergeInBetween(a, 3, 4, g))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var i int
	var end *ListNode
	cur, start := list1, list1
	for ; cur.Next != nil; i++ {
		if i == a-1 {
			start = cur
		}
		if i == b+1 {
			end = cur
		}
		cur = cur.Next
	}
	if end == nil {
		end = cur
	}
	cur = list2
	for cur.Next != nil {
		cur = cur.Next
	}
	start.Next = list2
	cur.Next = end
	return list1
}

func p(n *ListNode) {
	for ; n != nil; n = n.Next {
		fmt.Println(n.Val)
	}
}
