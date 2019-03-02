/*
给定两个二进制字符串，返回他们的和（用二进制表示）。
输入为非空字符串且只包含数字 1 和 0。

示例 1:
输入: a = "11", b = "1"
输出: "100"

示例 2:
输入: a = "1010", b = "1011"
输出: "10101"
*/
package main

import "fmt"

func addBinary(a string, b string) string {
	var carry, sum int
	i, j := len(a), len(b)
	if i < j {
		i, j = j, i
		a, b = b, a
	}

	res := make([]byte, i+1)

	for j > 0 {
		j--
		i--
		sum = int(a[i]-'0') + int(b[j]-'0') + carry
		carry = sum / 2
		sum = sum % 2
		res[i+1] = byte(sum + '0')
	}

	for i > 0 {
		i--
		sum = int(a[i]-'0') + carry
		carry = sum / 2
		sum = sum % 2
		res[i+1] = byte(sum + '0')
	}

	res[0] = byte(carry + '0')

	for i < len(res)-1 {
		if res[i] == '0' {
			i++
		} else {
			break
		}
	}
	return string(res[i:])
}

func main() {

}
