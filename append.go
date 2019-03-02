package main

import (
	"fmt"
)

func main() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4, 5, 6, 7)
	printSlice("a", a)

	var double = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, v := range double {
		fmt.Printf("2*%d = %d\n", v, v*2)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
