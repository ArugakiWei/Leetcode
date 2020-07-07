package main

import (
	"fmt"
)

func main() {
	//l1 := &ListNode{
	//	Val: 9,
	//	Next: &ListNode{
	//		Val: 1,
	//	},
	//}
	//l2 := &ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 1,
	//	},
	//}
	//addTwoNumbers(l1, l2)

	l3 := &ListNode{
		Val: 9,
	}
	l4 := &ListNode{
		Val: 5,
	}
	addTwoNumbers(l3, l4)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func getLength(l *ListNode) int {
	var length int
	for {
		if l == nil {
			break
		}
		length++

		if l.Next == nil {
			break
		}
		l = l.Next
	}
	return length
}

func padding(l *ListNode, max int) {
	var index int
	for {
		if l == nil {
			break
		}
		index++

		if index == max {
			break
		}

		if l.Next == nil {
			l.Next = &ListNode{
				Val: -1,
			}
		}

		l = l.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lenl1 := getLength(l1)
	lenl2 := getLength(l2)

	if lenl1 > lenl2 {
		padding(l2, lenl1)
	}
	if lenl1 < lenl2 {
		padding(l1, lenl2)
	}

	var head *ListNode

	tail := &ListNode{}
	var isCarry bool
	for {
		if l1 == nil || l2 == nil {
			if isCarry {
				tmp := &ListNode{
					Val: 1,
				}
				tail.Next = tmp
				tail = tmp
			}
			break
		}

		var val int
		if l1.Val == -1 && l2.Val != -1 {
			val = l2.Val
		}
		if l1.Val != -1 && l2.Val == -1 {
			val = l1.Val
		}
		if l1.Val != -1 && l2.Val != -1 {
			val = l1.Val + l2.Val
		}
		if l1.Val == -1 && l2.Val == -1 {
			break
		}
		if isCarry {
			val = val + 1
			isCarry = false
		}

		if val >= 10 {
			val = val - 10
			isCarry = true
		}

		tmp := &ListNode{
			Val: val,
		}
		if head == nil {
			head = tmp
		}
		tail.Next = tmp
		tail = tmp

		l1 = l1.Next
		l2 = l2.Next
	}
	return head
}

func show(l1 *ListNode) {
	for {
		fmt.Print(l1.Val)
		if l1 == nil {
			fmt.Println()
			break
		}
		if l1.Next == nil {
			fmt.Println()
			break
		}
		fmt.Print("->")
		l1 = l1.Next
	}
}