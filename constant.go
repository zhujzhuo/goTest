package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000
	const m = 1

	const d = 3e20 / n

	const e = 3e10 / m

	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(int64(e)) //如果int64放不下强制转化后的值，就会报错没法把float的值转化为int

	fmt.Println(math.Sin(n))
}
