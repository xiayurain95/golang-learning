package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([]Interval{{1, 3}, {2, 6}, {4, 10}, {15, 18}}))
}

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	sort.Sort(intercalSlice(intervals))
	for i := 0; i < len(intervals)-1; {
		j := i + 1
		if intervals[i].End >= intervals[j].Start {
			intervals[j].Start = intervals[i].Start
			if intervals[i].End > intervals[j].End {
				intervals[j].End = intervals[i].End
			}
			copy(intervals[i:], intervals[i+1:])
			intervals = intervals[:len(intervals)-1]
		} else {
			i++
		}
	}
	return intervals
}

type intercalSlice []Interval

func (i intercalSlice) Len() int           { return len(i) }
func (i intercalSlice) Less(j, k int) bool { return i[j].Start < i[k].Start }
func (i intercalSlice) Swap(j, k int)      { i[j], i[k] = i[k], i[j] }
