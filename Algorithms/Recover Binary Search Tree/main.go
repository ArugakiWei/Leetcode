package main

func main() {
	z := &TreeNode{
		Val: 2,
	}
	a := &TreeNode{
		Val: 1,
	}
	b := &TreeNode{
		Val:  4,
		Left: z,
	}
	c := &TreeNode{
		Val:   3,
		Left:  a,
		Right: b,
	}
	recoverTree(c)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	// 中序遍历
	nums := make([]int, 0, 0)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		nums = append(nums, root.Val)
		inorder(root.Right)
	}
	inorder(root)

	// 在中序遍历的结果中查询需要交换的2个值
	x, y := findTwoSwapped(nums)
	recover1(root, 2, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			y = nums[i+1]
			if x == -1 {
				x = nums[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func recover1(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover1(root.Right, count, x, y)
	recover1(root.Left, count, x, y)
}
