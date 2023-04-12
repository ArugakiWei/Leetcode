package main

import "fmt"

func main() {
	// a5 := &ListNode{
	// 	Val:  5,
	// 	Next: nil,
	// }
	// a4 := &ListNode{
	// 	Val:  4,
	// 	Next: a5,
	// }
	// a3 := &ListNode{
	// 	Val:  3,
	// 	Next: a4,
	// }
	a2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	a1 := &ListNode{
		Val:  1,
		Next: a2,
	}
	fmt.Println(removeNthFromEnd(a1, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	slow, fast, index := dummy, dummy, 0
	for fast != nil {
		if index > n {
			slow = slow.Next
		}
		fast = fast.Next
		index++
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
