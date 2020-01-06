package main

import (
	"fmt"
	"sort"
)

// 数据区间合并
// input [1,3] [2,6] [10,20] [1,7] [0,4] [7,9]  [22,30]
// output [0,9] [10,20] [22,30]
type Interval struct {
	Start int
	End   int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	// 先排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	newIts := make([]Interval, 0, len(intervals))
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= temp.End {
			temp.End = max(temp.End, intervals[i].End)
		} else {
			newIts = append(newIts, temp)
			temp = intervals[i]
		}
	}
	newIts = append(newIts, temp)

	return newIts
}

func main() {

	a := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {9, 20}}

	fmt.Println(merge(a))
}
