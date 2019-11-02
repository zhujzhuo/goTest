/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:
所有输入只包含小写字母 a-z 。
*/

package main

import "fmt"

//最长公共子串，首先取第一个字符串为基准，一个个的和剩余的子串对应位置比较，相同则继续，有一个不同，返回index，并且返回[:index]
func longestCommonPrefix(strs []string) string {
	//空strs的判断和只有一个子串的数组
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	var basestring string = strs[0]
	var strslen int = len(strs)
	for i := 0; i < len(basestring); i++ {
		tmpstr := basestring[i]
		for j := 1; j < strslen; j++ {
			if i == len(strs[j]) || strs[j][i] != tmpstr {  //在basestring达到最短字符串长度 或者 出现不相等字符时候 返回即可，否则一直比对下去
				return basestring[:i]
			}
		}
	}
	return basestring
}

func main() {
	var strs1 []string = []string{"dog", "racecar", "car"}
	var strs2 []string = []string{"flower", "flow", "flight"}
	var strs3 []string = []string{"flgggower", "flgggow", "flgggight"}
	var strs4 []string = []string{"aa", "a"}
	var strs5 []string = []string{"a", "a", "a"}
	var strs6 []string = []string{""}
	var strs7 []string = []string{"", ""}
	var strs8 []string = []string{"aa", "aa"}
	var strs9 []string = []string{"ca", "a"}
	var strs10 []string = []string{"a", "ca"}
	var strs11 []string = []string{}
	fmt.Println(longestCommonPrefix(strs1))
	fmt.Println(longestCommonPrefix(strs2))
	fmt.Println(longestCommonPrefix(strs3))
	fmt.Println(longestCommonPrefix(strs4))
	fmt.Println(longestCommonPrefix(strs5))
	fmt.Println(longestCommonPrefix(strs6))
	fmt.Println(longestCommonPrefix(strs7))
	fmt.Println(longestCommonPrefix(strs8))
	fmt.Println(longestCommonPrefix(strs9))
	fmt.Println(longestCommonPrefix(strs10))
	fmt.Println(longestCommonPrefix(strs11))
}
