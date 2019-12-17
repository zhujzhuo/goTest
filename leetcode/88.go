package main

import (
	"fmt"
	"strconv"
)

func add(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	result := ""
	var maxlen int
	if len1 >= len2 {
		maxlen = len1
	} else {
		maxlen = len2
	}
	mid := ""
	//短字符串补全高位为0
	if len1 >= len2 {
		for j := len1 - len2; j > 0; j-- {
			mid += "0"
		}
		str2 = mid + str2
	} else {
		for j := len2 - len1; j > 0; j-- {
			mid += "0"
		}
		str1 = mid + str1
	}
	carry := 0
	for i := maxlen - 1; i >= 0; i-- {
		int1, _ := strconv.Atoi(string(str1[i]))
		int2, _ := strconv.Atoi(string(str2[i]))
		num := (int1 + int2 + carry) % 10  //计算加和个位数
		carry = (int1 + int2 + carry) / 10 //计算加和进位
		result = strconv.Itoa(num) + result
		if i == 0 && carry == 1 {
			//最后一位进位还是1，需要加到结果集中
			result = "1" + result
		}
	}
	return result
}

func main() {
	var str1 string = "ksksks"
	var len1 int = len(str1)
	fmt.Println(len1)

	var str2 string = "99"
	var str3 string = "9999"
	fmt.Println(add(str2, str3))
	var str4 string = "10"
	var str5 string = "9999"
	fmt.Println(add(str4, str5))

}
