package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum1(head *ListNode) int {
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	var max int
	for i := 0; i <= len(vals)/2-1; i++ {
		v := vals[i] + vals[len(vals)-1-i]
		if max < v {
			max = v
		}
	}
	return max
}

// 快慢指针找中点
// 反转后半链表
func pairSum(head *ListNode) int {
	mid, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		mid = mid.Next
	}

	var prev, cur *ListNode
	cur = mid
	for cur != nil {
		t := cur.Next
		cur.Next = prev
		prev = cur
		cur = t
	}

	var max int
	for prev != nil {
		v := prev.Val + head.Val
		if v > max {
			max = v
		}
		prev = prev.Next
		head = head.Next
	}
	return max
}
