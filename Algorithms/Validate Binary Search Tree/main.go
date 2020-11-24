package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func isValidBST(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//
//	nums := make([]int, 0, 0)
//	var inorder func(root *TreeNode)
//	inorder = func(root *TreeNode) {
//		if root == nil {
//			return
//		}
//
//		inorder(root.Left)
//		nums = append(nums, root.Val)
//		inorder(root.Right)
//	}
//
//	inorder(root)
//
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] >= nums[i+1] {
//			return false
//		}
//	}
//	return true
//}

func isValidBST(root *TreeNode) bool {
	return isValidBST1(root, nil, nil)
}

func isValidBST1(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValidBST1(root.Left, min, root) && isValidBST1(root.Right, root, max)
}
