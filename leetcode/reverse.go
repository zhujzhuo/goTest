/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
示例 1:
输入: 123
输出: 321

示例 2:
输入: -123
输出: -321

示例 3:
输入: 1200
输出: 21
注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0
*/
package main

import "fmt"

var maxnum int = 1<<31 - 1
var minnum int = -(1 << 31)

func reverse(x int) int {
	var num int
	for x != 0 {
		num = num*10 + x%10 //如果x是负数，这里全部为负值的相加
		x = x / 10
	}
	if num > maxnum || num < minnum {
		return 0
	}
	return num
}

/*
import "math"
func reverse(x int) int {
	var r int
	var limit int
	if x >= 0 {
		limit = math.MaxInt32 / 10
		for x != 0 {
			remainder := x % 10
			if r > limit || (r == limit && remainder > 7) {
				return 0
			}
			r = r*10 + remainder
			x /= 10
		}
	} else {
		limit = math.MinInt32 / 10
		for x != 0 {
			remainder := x % 10
			if r < limit || (r == limit && remainder < -8) {
				return 0
			}
			r = r*10 + remainder
			x /= 10
		}
	}

	return r
}
*/
func main() {
	/*
	   fmt.Println(1<<31-1)
	   fmt.Println(1<<2)
	   var j int = 2
	   for i:=2;i<32;i++ {
	       j *= 2
	   }
	   fmt.Println(j-1)
	   fmt.Println(minnum)
	*/

	fmt.Println(maxnum)
	fmt.Println(minnum)
	fmt.Println(reverse(12))
	fmt.Println(reverse(987654321))
	fmt.Println(reverse(98765))
	fmt.Println(reverse(120000))
	fmt.Println(reverse(2147483648))
	fmt.Println(reverse(-123))

}
