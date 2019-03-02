/*
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
如果不存在最后一个单词，请返回 0 。
说明：一个单词是指由字母组成，但不包含任何空格的字符串。
示例:
输入: "Hello World"
输出: 5
*/

package main

import "fmt"

//去掉字符串末尾的空格，之后倒叙找到第一个空格位置，之后的单词就是最后一个单词，得到其长度即可
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	var start int = len(s) - 1
	var end int = len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		//去掉末尾的空格，直到有第一个不为空格的字符出现，全是空格最后也是会返回0
		if s[i] == ' ' && start == end {
			start--
			end--
		}
		//去掉末尾的空格之后，开始计算不为空格的字符个数
                //start的值=0 是解决"a"  "aa "这种字符串的
		if s[i] != ' ' && start >= 0 {
			start--
		}
		//直到第一个空格出现，此时最后一个单词的开始和结束是不相等的，最后返回end和start的差值
		if s[i] == ' ' && start != end {
			break
		}
	}
	return end - start

}

func main() {
	var str1 string = "Hello World"
	var str2 string = ""
	var str3 string = "Hello     World    cd"
	var str4 string = "    Hello     World    cd  "
	var str5 string = " "
	var str6 string = "aa"
	var str7 string = "a  "

	fmt.Println(lengthOfLastWord(str1))
	fmt.Println(lengthOfLastWord(str2))
	fmt.Println(lengthOfLastWord(str3))
	fmt.Println(lengthOfLastWord(str4))
	fmt.Println(lengthOfLastWord(str5))
	fmt.Println(lengthOfLastWord(str6))
	fmt.Println(lengthOfLastWord(str7))
}
