package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var out []int
	var solve func(node *TreeNode)
	solve = func(node *TreeNode) {
		if node.Left != nil {
			solve(node.Left)
		}
		out = append(out, node.Val)
		if node.Right != nil {
			solve(node.Right)
		}
	}
	if root != nil {
		solve(root)
	}
	return out
}
