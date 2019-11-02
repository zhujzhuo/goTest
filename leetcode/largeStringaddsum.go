package main

//错误的
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
	var carry int = 0
	mid := ""
	//短字符串补全高位为0
	if len1 >= len2 {
		for j := len1 - len2; j >= 0; j-- {
			mid += "0"
		}
		str2 += mid
	} else {
		for j := len1 - len2; j >= 0; j-- {
			mid += "0"
		}
		str1 += mid
	}

	for i := maxlen - 1; i >= 0; i-- {
		num := (strconv.Atoi(str1[i]) + strconv.Atoi(str2[i]) + carry) % 10   //计算加和个位数
		carry := (strconv.Atoi(str1[i]) + strconv.Atoi(str2[i]) + carry) / 10 //计算加和进位
		result += strconv.Itoa(num)
		if i == 0 && carry == 1 { //最后一位进位还是1，需要加到结果集中
			result += "1"
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

}
