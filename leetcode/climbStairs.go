/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

示例 1：
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶

示例 2：
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

4
1 1 1 1
1 2 1
2 2
1 1 2
2 1 1
5种
*/

package main

import "fmt"

/*
这是一个fab数列，从n=1开始
fab
1
2
3
5
8
*/

/*
// 如下递归正确性没什么问题，但是时间上不满足条件
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
*/

//直接用循环的方法做，执行时间上消耗时间较少
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		var s1 int = 1
		var s2 int = 2
		var si int = 0
		for i := 3; i <= n; i++ {
			si = s1 + s2
			s1 = s2
			s2 = si
		}
		return si

	}
}

func main() {
	var num1 int = 0
	var num2 int = 1
	var num3 int = 3
	var num4 int = 4
	var num5 int = 5
	fmt.Println(climbStairs(num1))
	fmt.Println(climbStairs(num2))
	fmt.Println(climbStairs(num3))
	fmt.Println(climbStairs(num4))
	fmt.Println(climbStairs(num5))
}
