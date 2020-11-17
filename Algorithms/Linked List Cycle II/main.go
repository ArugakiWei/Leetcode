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
	fmt.Println(detectCycle(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	flag := false

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			flag = true
			break
		}
	}

	if flag {
		slow = head
		for slow != fast {
			slow = slow.Next
			fast = fast.Next
		}
		return slow
	}
	return nil
}
