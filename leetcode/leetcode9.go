/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:
输入: 121
输出: true

示例 2:
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

示例 3:
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
*/

package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x <= 9 {
		return true
	} else if x%10 == 0 {
		return false
	}
	var num int
	//传入的x在下面的计算中会被重新赋值，需要使用中间变量来保存x的初始值
	var xtmp int = x
	for x != 0 {
		// 回文计算：乘以10+余数
		num = num*10 + x%10
		x = x / 10
	}
	if xtmp == num {
		return true
	} else {
		return false
	}
}

func main() {
	var i, j, k, l int = 121, 2333, -123, 10
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(i))
	fmt.Println(isPalindrome(j))
	fmt.Println(isPalindrome(k))
	fmt.Println(isPalindrome(l))
	fmt.Println(isPalindrome(0))
}
