package main

import "fmt"

func main() {
	c := &ListNode{
		Val: 4,
	}
	b := &ListNode{
		Val:  0,
		Next: c,
	}
	a := &ListNode{
		Val:  2,
		Next: b,
	}
	c.Next = a
	head := &ListNode{
		Val:  3,
		Next: a,
	}
	fmt.Println(hasCycle(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}
