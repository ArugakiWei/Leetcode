package main

import (
	"fmt"
)

func main() {
	a := &ListNode{
		Val: 3,
	}
	b := &ListNode{
		Val:  2,
		Next: a,
	}
	c := &ListNode{
		Val:  1,
		Next: b,
	}

	output(c)
	output(reverseList(c))
}

func output(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print("->")
		}
		head = head.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre, cur *ListNode
	pre = head
	for pre != nil {
		tmp := pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
	}

	return cur
}
