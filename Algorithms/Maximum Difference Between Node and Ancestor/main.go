package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	var val func(n *TreeNode) (int, int)
	val = func(n *TreeNode) (int, int) {
		if n == nil {
			return 100001, -1
		}
		minLeft, maxLeft := val(n.Left)
		minRight, maxRight := val(n.Right)
		if n.Left != nil || n.Right != nil {
			ans = max(ans, abs(n.Val-min(minLeft, minRight)))
			ans = max(ans, abs(n.Val-max(maxLeft, maxRight)))
		}
		return min3(minLeft, minRight, n.Val), max3(maxLeft, maxRight, n.Val)
	}
	val(root)
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}
