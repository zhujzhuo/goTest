package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

//i  := 2   这里的赋值编译的时候会报错，函数外不可使用这种赋值语句
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("Now you have %f problems.", math.Nextafter(2, 3))
	fmt.Println(math.Pi)
	fmt.Println(math.MaxUint8)
	fmt.Println(add(42, 13))
	fmt.Println(split(17))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	var i int
	fmt.Println(i, c, python, java)
	k := 3
	//在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义
	//函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
	fmt.Println(k)
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	var x, y int = 3, 4
	var t float64 = math.Sqrt(float64(x*x + y*y))
	var tt int = int(t)
	fmt.Println(x, y, tt)

	const Pi = 3.14
	//常量不能使用 := 语法定义。

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	//fmt.Println(needInt(Big))    constant 1267650600228229401496703205376 overflows int
}
