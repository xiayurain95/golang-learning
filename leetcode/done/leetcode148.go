package main

import (
	"sort"
)

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := make([]*ListNode, 0)
	for head != nil {
		tmp = append(tmp, head)
		head = head.Next
	}
	sort.Sort(ListNodeSlice(tmp))
	for i := 0; i < len(tmp)-1; i++ {
		tmp[i].Next = tmp[i+1]
	}
	tmp[len(tmp)-1].Next = nil
	return tmp[0]
}

type ListNodeSlice []*ListNode

func (l ListNodeSlice) Len() int           { return len(l) }
func (l ListNodeSlice) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ListNodeSlice) Less(i, j int) bool { return l[i].Val < l[j].Val }
