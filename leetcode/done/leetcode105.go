package main

func main() {

}
func buildTree(preorder []int, inorder []int) *TreeNode {
	var solve func(order []int) *TreeNode
	solve = func(order []int) *TreeNode {
		if len(preorder) == 0 || len(order) == 0 {
			return nil
		}
		i := 0
		for ; i < len(order); i++ {
			if preorder[0] == order[i] {
				break
			}
		}
		node := new(TreeNode)
		node.Val = preorder[0]
		preorder = preorder[1:]
		node.Left = solve(order[:i])
		node.Right = solve(order[i+1:])
		return node
	}
	return solve(inorder)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
