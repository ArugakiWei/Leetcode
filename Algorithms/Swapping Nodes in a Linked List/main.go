package main

import "fmt"

func main() {
	a5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	a4 := &ListNode{
		Val:  4,
		Next: a5,
	}
	a3 := &ListNode{
		Val:  3,
		Next: a4,
	}
	a2 := &ListNode{
		Val:  2,
		Next: a3,
	}
	a1 := &ListNode{
		Val:  1,
		Next: a2,
	}
	fmt.Println(swapNodes(a1, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	// prepk 第k-1个节点, prepnk 第n-k-1个节点
	// pk 第k个节点, pnk 第n-k个节点
	// 找到上述2组4个节点, 互换其 next 指针即可
	p, prepk, prepnk, pk, pnk := dummy, dummy, dummy, dummy, dummy
	for i := 0; i < k; i++ {
		if i < k-1 {
			prepk = prepk.Next
		}
		p = p.Next
	}
	pk = p
	for p != nil {
		if p.Next == nil {
			prepnk = pnk
		}
		p = p.Next
		pnk = pnk.Next
	}
	prepk.Next, prepnk.Next = prepnk.Next, prepk.Next
	pk.Next, pnk.Next = pnk.Next, pk.Next
	return dummy.Next
}
