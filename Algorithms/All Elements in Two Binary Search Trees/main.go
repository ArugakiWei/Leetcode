package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	ans1 := inorder(root1)
	ans2 := inorder(root2)
	return merge(ans1, ans2)
}

func merge(a1, a2 []int) []int {
	var ans []int
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			ans = append(ans, a1[i])
			i++
		} else {
			ans = append(ans, a2[j])
			j++
		}
	}
	if len(a1) > i {
		ans = append(ans, a1[i:]...)
	} else {
		ans = append(ans, a2[j:]...)
	}
	return ans
}

func inorder(node *TreeNode) []int {
	var ans []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(node)
	return ans
}
