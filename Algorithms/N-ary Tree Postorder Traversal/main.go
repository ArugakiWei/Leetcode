package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var ans []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, n := range node.Children {
			dfs(n)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return ans
}
