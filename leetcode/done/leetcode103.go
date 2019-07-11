package main

func main() {
	//zigzagLevelOrder()
}
func zigzagLevelOrder(root *TreeNode) [][]int {
	var out [][]int
	if root == nil {
		return out
	}

	ch := make(chan *TreeNode, 1000)
	ch <- root
	ch <- nil
	var tmp []int
	var pre *TreeNode
	for len(ch) != 0 {
		node := <-ch
		if node == nil && pre != nil {
			out = append(out, tmp)
			tmp = make([]int, 0)
			ch <- nil
		} else if node == nil && pre == nil {
		} else {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				ch <- node.Left
			}
			if node.Right != nil {
				ch <- node.Right
			}
		}
		pre = node
	}
	for i := 0; i < len(out); i++ {
		if i%2 == 1 {
			for j := 0; j < len(out[i])/2; j++ {
				out[i][j], out[i][len(out[i])-1-j] = out[i][len(out[i])-1-j], out[i][j]
			}
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
