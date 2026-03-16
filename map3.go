package main

import "fmt"

func main() {

	m := make(map[string]int, 10)
	n := make([]int, 2, 10)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))
	fmt.Println("slice len:", len(n))
	fmt.Println("slice cap:", cap(n))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	l := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", l)
}
