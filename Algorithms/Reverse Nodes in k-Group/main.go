package main

import (
	"fmt"
	"time"
)

func main() {
	e := &ListNode{
		Val: 5,
	}
	d := &ListNode{
		Val:  4,
		Next: e,
	}
	a := &ListNode{
		Val:  3,
		Next: d,
	}
	b := &ListNode{
		Val:  2,
		Next: a,
	}
	c := &ListNode{
		Val:  1,
		Next: b,
	}

	output(reverseKGroup(c, 2))
}

func output(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print("->")
		}
		head = head.Next
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	length := 0

	t := head
	for t != nil {
		length++
		t = t.Next
	}

	start, end := 1, k
	res := head
	for {
		res = reverseBetween(res, start, end)
		output(res)
		start, end = start+k, end+k
		if end > length {
			break
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var cur, pre, rHead, rEnd *ListNode

	i, rH, rE := 1, 0, 0
	pre = head

	for pre != nil {
		if i > n {
			break
		}

		if i < m {
			i++
			if cur == nil {
				cur = head
			} else {
				cur = cur.Next
			}
			pre = pre.Next
			continue
		}

		if rH == 0 {
			rH++
			rHead = cur
		}
		if rE == 0 {
			rE++
			rEnd = pre
		}

		i++
		tmp := pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
	}

	if rHead != nil {
		rHead.Next = cur
	}
	if rEnd != nil {
		rEnd.Next = pre
	}

	if rHead == nil {
		return cur
	}
	return head
}
