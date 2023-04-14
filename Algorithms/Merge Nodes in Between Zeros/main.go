package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes1(head *ListNode) *ListNode {
	dummy := &ListNode{}
	p, p1 := head.Next, dummy
	count := 0
	for p != nil {
		if p.Val == 0 {
			p1.Next = &ListNode{
				Val: count,
			}
			p1 = p1.Next
			count = 0
		}
		count += p.Val
		p = p.Next
	}
	return dummy.Next
}

func mergeNodes(head *ListNode) *ListNode {
	p, p1 := head.Next, &ListNode{}
	dummy := p1
	count := 0
	for p != nil {
		if p.Val == 0 {
			p.Val = count
			p1.Next = p
			p1 = p
			count = 0
		} else {
			count += p.Val
		}
		p = p.Next
	}
	return dummy.Next
}
