package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 6)
	a = append(a, 7, 8)
	b := []int{9, 10, 11}
	a = append(a, b...)
	sum(a...)
}
