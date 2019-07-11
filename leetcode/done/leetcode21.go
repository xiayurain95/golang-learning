package main

func main() {

}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := new(ListNode)
	out := p

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return out.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
