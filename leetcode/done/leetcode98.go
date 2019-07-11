package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var inOrder func(*TreeNode)
	var out []int
	inOrder = func(node *TreeNode) {
		if node != nil {
			inOrder(node.Left)
			out = append(out, node.Val)
			inOrder(node.Right)
		}
	}
	inOrder(root)
	for i := 0; i < len(out)-1; i++ {
		if out[i] > out[i+1] {
			return false
		}
	}
	return true
}
