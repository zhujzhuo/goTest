/*
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
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
*/
package main

import (
	"fmt"
	"math"
)
//因此需要先牢记这样的规则:小写字母开头的函数只在本包内可见，大写字母开头的函数才能被其他包使用
//函数的闭包
//闭包是一个函数值，它来自函数体的外部的变量引用。 
//函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
                        i,
			pos(i),
			neg(-2*i),
		)
	}
}

