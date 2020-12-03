package main

import "fmt"

func main() {
	d := &ListNode{
		Val: 1,
	}
	a := &ListNode{
		Val:  2,
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
	fmt.Println(isPalindrome(c))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	slow = reverse(slow)

	for slow != nil {
		if slow.Val != head.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	var cur, next *ListNode
	next = node
	for next != nil {
		t := next.Next
		next.Next = cur
		cur = next
		next = t
	}
	return cur
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
