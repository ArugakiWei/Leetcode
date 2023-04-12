package main

import (
	"fmt"
	"time"
)

func main() {
	a6 := &ListNode{
		Val:  2,
		Next: nil,
	}
	a5 := &ListNode{
		Val:  5,
		Next: a6,
	}
	a4 := &ListNode{
		Val:  2,
		Next: a5,
	}
	a3 := &ListNode{
		Val:  3,
		Next: a4,
	}
	a2 := &ListNode{
		Val:  4,
		Next: a3,
	}
	a1 := &ListNode{
		Val:  1,
		Next: a2,
	}
	// log(partition(a1, 1))
	log(partition(a1, 3))
	// log(partition(a1, 6))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	// 存放小于 x 的链表的虚拟头结点
	dummy1 := &ListNode{-1, nil}
	// 存放大于等于 x 的链表的虚拟头结点
	dummy2 := &ListNode{-1, nil}
	// p1, p2 指针负责生成结果链表
	p1, p2 := dummy1, dummy2
	// p 负责遍历原链表，类似合并两个有序链表的逻辑
	// 这里是将一个链表分解成两个链表
	p := head
	for p != nil {
		if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		// 断开原链表中的每个节点的 next 指针
		temp := p.Next
		// p.Next = nil
		p = temp
	}
	log(dummy1.Next)
	log(dummy2.Next)
	// 连接两个链表
	p1.Next = dummy2.Next

	return dummy1.Next
}

func log(n *ListNode) {
	for n != nil {
		fmt.Printf("%d->", n.Val)
		n = n.Next
		time.Sleep(time.Second)
	}
	fmt.Println()
}
