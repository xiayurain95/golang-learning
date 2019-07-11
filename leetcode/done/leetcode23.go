package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else {
		tmp1, tmp2 := lists[len(lists)-1], lists[len(lists)-2]
		p := new(ListNode)
		out := p
		for tmp1 != nil && tmp2 != nil {
			if tmp1.Val < tmp2.Val {
				p.Next = tmp1
				tmp1 = tmp1.Next
			} else {
				p.Next = tmp2
				tmp2 = tmp2.Next
			}
			p = p.Next
		}
		if tmp1 != nil {
			p.Next = tmp1
		}
		if tmp2 != nil {
			p.Next = tmp2
		}
		return mergeKLists(append(lists[:len(lists)-2], out.Next))
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
