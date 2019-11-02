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
     var  a  string = "1011"
     var  b  string = "1011"
     var  c  string = "1000"
     fmt.Println(addBinary(a,b))   //10110
     fmt.Println(addBinary(a,c))   //10011
}

/*

func addBinary(a string, b string) string {
    result := make([]byte, 0, len(a))
    var addi byte = '0'
    i := len(a)-1
    j := len(b)-1
    for ; i >= 0 && j >= 0; {
        c := a[i]+b[j]+addi-'0'-'0'
        if c >= '2'{
            c = c - '2' + '0'
            addi = '1'
        } else {
            addi = '0'
        }
        result = append(result, c)
        i--
        j--
    }
    for i >= 0 {
        c := a[i] + addi -'0'
        if c >= '2'{
            c = c - '2' + '0'
            addi = '1'
        } else {
            addi = '0'
        }
        result = append(result, c)
        i--
    }
    for j >= 0 {
        c := b[j] + addi -'0'
        if c >= '2'{
            c = c - '2' + '0'
            addi = '1'
        } else {
            addi = '0'
        }
        result = append(result, c)
        j--
    }
    if addi > '0' {
        result = append(result, addi)
    }
    r := make([]byte, 0, len(result))
    for index := len(result)-1; index >= 0; index-- {
        r = append(r, result[index])
    }
    return string(r)
}
    
*/
