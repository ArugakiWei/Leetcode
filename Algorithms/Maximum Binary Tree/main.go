package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	val, index := max(nums)
	return &TreeNode{
		Val:   val,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}

func max(nums []int) (int, int) {
	ans, index := 0, 0
	for i, v := range nums {
		if v > ans {
			ans, index = v, i
		}
	}
	return ans, index
}
