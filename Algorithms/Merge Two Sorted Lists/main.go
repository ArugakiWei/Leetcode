package main

import (
	"fmt"
	"time"
)

func main() {
	l13 := &ListNode{Val: 4}
	l12 := &ListNode{Val: 2, Next: l13}
	l11 := &ListNode{Val: 1, Next: l12}
	l23 := &ListNode{Val: 4}
	l22 := &ListNode{Val: 3, Next: l23}
	l21 := &ListNode{Val: 1, Next: l22}
	output(mergeTwoLists(l11, l21))
}

func output(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1 == nil && l2 == nil {
		return nil
	}

	var res *ListNode
	if l1.Val < l2.Val {
		res = &ListNode{
			Val: l1.Val,
		}
		l1 = l1.Next
	} else {
		res = &ListNode{
			Val: l2.Val,
		}
		l2 = l2.Next
	}

	head := res
	for l1 != nil && l2 != nil {
		min := 0
		if l1.Val < l2.Val {
			min = l1.Val
			l1 = l1.Next
		} else {
			min = l2.Val
			l2 = l2.Next
		}
		newNode := &ListNode{
			Val: min,
		}
		head.Next = newNode
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}

	return res
}
