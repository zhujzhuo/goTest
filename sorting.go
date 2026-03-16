package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {

	//Go’s sort package implements sorting for builtins and user-defined types. We’ll look at sorting for builtins first.
	//Sort methods are specific to the builtin type; here’s an example for strings. Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.

	strs := []string{"ssc", "all", "bc"}
	//排序可以使用slices的Sort或者sort的Strings
	slices.Sort(strs)
	//sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
