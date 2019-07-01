package code

import (
	"fmt"
	"sort"
)

func SortInp() {
	var a sortEx = []int{1, 34, 4, 2, 35, 2, 45, 2, 2}
	sort.Sort(a)
	fmt.Printf("%v", a)
}

type sortEx []int

func (ins sortEx) Less(i, j int) bool {
	if ins[i] < ins[j] {
		return true
	} else {
		return false
	}
}
func (ins sortEx) Swap(i, j int) { ins[i], ins[j] = ins[j], ins[i] }
func (ins sortEx) Len() int      { return len(ins) }
