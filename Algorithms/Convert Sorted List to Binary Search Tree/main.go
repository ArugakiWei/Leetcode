package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := findRoot(left, right)
	return &TreeNode{
		Val:   mid.Val,
		Left:  buildTree(left, mid),
		Right: buildTree(mid.Next, right),
	}
}

func findRoot(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right && fast.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
