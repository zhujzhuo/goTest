/*
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1

说明:
当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	var hs_ln int = len(haystack)
	var nd_ln int = len(needle)
	if hs_ln < nd_ln {
		return -1
	}
        //字符串包含判断，判断i到i+子串长度的串是否和子串相等
	for i := 0; i <= hs_ln-nd_ln; i++ { 
		if haystack[i:i+nd_ln] == needle {
			return i
		}
	}
        return -1
}
func main() {
	var haystack = "hello"
	var needle = "ll"
	fmt.Println(strStr(haystack, needle))

        var str1 string = "hel"
        var str2 string = ""
        fmt.Println(strStr(str1,str2))
        
        var str3 string = "helhytg"
        var str4 string = "hytg"
        fmt.Println(strStr(str3,str4))
}
