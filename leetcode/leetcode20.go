/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true
*/

package main

import (
	"fmt"
)

//堆栈的push pop
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	// 在go中 '[' '}' ')' 等不能按照string处理，需要按照byte处理,使用单引号
	// 三种初始化map的方法
	//m := make(map[byte]byte)
	//m['('] = ')'
	//m['['] = ']'
	//m['{'] = '}'
	var m = map[byte]byte{'(': ')', '{': '}', '[': ']'}
	//m := map[byte]byte{'(':')','{':'}','[':']'}
	sli := make([]byte, 0)
	var flag bool = true
	//遍历string s
	// 类似遇到前括弧则入栈，遇到后括弧则和栈顶元素比对
	for i, b := range s {
		if b == '[' || b == '(' || b == '{' {
			sli = append(sli, byte(b))
		} else if (b == ']' || b == ')' || b == '}') && len(sli) > 0 { // len判断栈为空时进来了后括弧
			tmpbyte := sli[len(sli)-1]
			sli = sli[:len(sli)-1] // sli 减掉最后一个byte
			if m[tmpbyte] != byte(b) {
				flag = false
				return flag
			}
		} else if (b == ']' || b == ')' || b == '}') && len(sli) == 0 {
			flag = false
			return flag

		}
		if i == len(s)-1 && len(sli) != 0 { //最后一个元素比对完，但是sli还有前括号则返回false
			flag = false
		}
	}
	return flag
}

func main() {
	var str1 string = ""
	var str2 string = "()"
	var str3 string = "(){}[]"
	var str4 string = "{[()]}"
	var str5 string = "[()]}"
	var str6 string = "}}"

	fmt.Println(isValid(str1))
	fmt.Println(isValid(str2))
	fmt.Println(isValid(str3))
	fmt.Println(isValid(str4))
	fmt.Println(isValid(str5))
	fmt.Println(isValid(str6))

}
