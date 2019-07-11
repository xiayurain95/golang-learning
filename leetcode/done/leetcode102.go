package main

func main() {

}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	m := map[int][]int{}
	var out [][]int
	var solve func(*TreeNode, int)
	solve = func(node *TreeNode, i int) {
		if node != nil {
			solve(node.Left, i+1)
			if _, ok := m[i]; !ok {
				m[i] = make([]int, 0)
			}
			m[i] = append(m[i], node.Val)
			solve(node.Right, i+1)
		}
	}
	solve(root, 0)
	for i := 0; i < len(m); i++ {
		out = append(out, m[i])
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
