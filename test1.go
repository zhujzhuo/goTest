package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4)) //5

	fmt.Println(compute(hypot))    // 5
	fmt.Println(compute(math.Pow)) //math.Pow(x, y), x的y次方   3*3*3*3=81
}
