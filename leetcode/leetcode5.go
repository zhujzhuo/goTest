/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"
*/
package main
 
import (
	"fmt"
)
 
// 中心扩展法
func LookPanlind_two(str string) (s string) {

    i, j, max, c := 0, 0, 0, 0
    for i = 0; i < len(str); i++ {
        // 回文长度为奇数
        for j = 0; (i+j) < len(str) && (i-j) >= 0; j++ {
            if str[i-j] != str[i+j] {
                break
            }

            c = 2*j + 1
        }
        if max < c {
            s = str[(i - j + 1):(i + j)]
            // println("奇数", s)
            max = c
        }
        // 回文长度为偶数
        for j = 0; (i+j+1) < len(str) && (i-j) >= 0; j++ {
            if str[i-j] != str[i+j+1] {
                break
            }

            c = 2*j + 2
        }
        if max < c {
            s = str[(i - j + 1):(i + j + 1)]
            // println("偶数", s)
            max = c
        }
    }
    return s
}


func   main(){

	var  str1 string= "abcddcba"
	var  str2 string= "abcdcba"
	var  str3 string = "abcsssasssddcba"
 

	fmt.Println(LookPanlind_two(str1))
	fmt.Println(LookPanlind_two(str2))
	fmt.Println(LookPanlind_two(str3))
}