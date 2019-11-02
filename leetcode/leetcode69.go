/*
实现 int sqrt(int x) 函数。
计算并返回 x 的平方根，其中 x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:
输入: 4
输出: 2

示例 2:
输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
由于返回类型是整数，小数部分将被舍去。
*/
package main 

import (
"fmt"
"math"
)

//这里可以不适用库函数完成，后面再做进一步的解题
//math的函数返回和传入参数基本都是float64类型，这里要做int和float64之间的转化
func mySqrt(x int) int {
     xtmp := float64(x)
     return int(math.Floor(math.Sqrt(xtmp)))
}

func main(){
     var num1 int = 0
     var num2 int = 4
     var num3 int = 8
     var num4 int = 9
     var num5 int = 10
     fmt.Println(mySqrt(num1))
     fmt.Println(mySqrt(num2))
     fmt.Println(mySqrt(num3))
     fmt.Println(mySqrt(num4))
     fmt.Println(mySqrt(num5))
}

/*
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	r := x
        //r^r 大于x，则重新赋值r， r = (r+x/r)/2  =>  2r=r+x/r => r=x/r
	for ; r > x/r; {
		r = (r+x/r)/2
	}

	return r
}
*/
