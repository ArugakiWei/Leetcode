package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	ListNode *ListNode
	Index    int
}

func main() {

}

func mergeKLists(lists []*ListNode) *ListNode {
	h := NewHeap(len(lists))
	for i := range lists {
		if lists[i] != nil {
			h.insert(&Node{
				ListNode: lists[i],
				Index:    i,
			})
			lists[i] = lists[i].Next
		}
	}
	dummy := &ListNode{Val: -1}
	p := dummy
	for h.len() != 0 {
		v := h.delMin()
		p.Next = v.ListNode
		p = p.Next

		if lists[v.Index] != nil {
			h.insert(&Node{
				ListNode: lists[v.Index],
				Index:    v.Index,
			})
			lists[v.Index] = lists[v.Index].Next
		}
	}
	return dummy.Next
}

type Heap struct {
	Data []*Node
	Size int
}

func NewHeap(cap int) *Heap {
	return &Heap{
		Data: make([]*Node, cap+1),
	}
}

func (h *Heap) print() {
	for _, v := range h.Data {
		if v != nil {
			fmt.Print(v.ListNode.Val)
			fmt.Print(",")
		}
	}
	fmt.Println()
}

func (h *Heap) parent(x int) int {
	return x / 2
}

func (h *Heap) right(x int) int {
	return x*2 + 1
}

func (h *Heap) left(x int) int {
	return x * 2
}

func (h *Heap) len() int {
	return h.Size
}

func (h *Heap) swim(x int) {
	for x > 1 && h.less(x, h.parent(x)) {
		h.swap(h.parent(x), x)
		x = h.parent(x)
	}
	return
}

func (h *Heap) sink(x int) {
	for h.left(x) <= h.len() {
		min := h.left(x)
		if h.right(x) <= h.len() && h.less(h.right(x), min) {
			min = h.right(x)
		}
		if h.less(x, min) {
			break
		}
		h.swap(x, min)
		x = min
	}
	return
}

func (h *Heap) delMin() *Node {
	min := h.Data[1]
	h.swap(1, h.Size)
	h.Data[h.Size] = nil
	h.Size--
	h.sink(1)
	return min
}

func (h *Heap) insert(n *Node) {
	h.Size++
	h.Data[h.Size] = n
	h.swim(h.Size)
}

func (h *Heap) less(x, y int) bool {
	return h.Data[x].ListNode.Val < h.Data[y].ListNode.Val
}

func (h *Heap) swap(x, y int) {
	h.Data[x], h.Data[y] = h.Data[y], h.Data[x]
}
