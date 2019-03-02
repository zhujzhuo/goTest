/*
如果有非7为的ASCII码的字符在字符串中，需要转化为rune的切片，再按照索引输出，输出的是字符而不是字节，否则多字节的字符按照原始输出将会是单个的字节值，会有问题，并且输出时转化为string(char)将输出string，否则输出的是ASCII的编码code
chars := []rune(text1)
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var str = "hello 你好"

	//golang中string底层是通过byte数组实现的，索引直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str)) //12   6+6

	//以下两种都可以得到str的字符串长度
	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str)) //8

	//通过rune类型处理unicode字符
	// 将Go语言的字符串转化为Unicode码点切片（类型为 [ ]rune），切片是支持直接索引的。
	fmt.Println("rune:", len([]rune(str))) //8
	//
	var s string = "MCMXCIV"
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r) //  0 M 77 .etc
	}
	fmt.Println(string(s[2])) //M   强制转化byte到string
	fmt.Println(s[2])         //77  byte
	fmt.Println(s[2:])        //MXCIV

	text1 := "abcdefg在"
	fmt.Println(text1[1])                      //获取字符串索引位置为n的原始字节，比如a为97
	fmt.Println(text1[1:3])                    //截取得字符串索引位置为 n 到 m-1 的字符串
	fmt.Println(text1[1:])                     //截取得字符串索引位置为 n 到 len(s)-1 的字符串
	fmt.Println(text1[:3])                     //截取得字符串索引位置为 0 到 m-1 的字符串
	fmt.Println(len(text1))                    //获取字符串的字节数  10
	fmt.Println(utf8.RuneCountInString(text1)) //获取字符串字符的个数  8
	fmt.Println([]rune(text1))                 // 将字符串的每一个字节转换为码点值，比如这里会输出[97 98 99 100 101 102 103 22312]
	fmt.Println(string(text1[1]))              // 获取字符串索引位置为n的字符值
	for index, char := range text1 {
		fmt.Printf("%d %U %c \n", index, char, char)
	}
        // 有些字符可能有多个字节,
        fmt.Println("================")
        fmt.Println(string(text1[7])) //  这个地方输出会有问题，"在" 这里有三个字节，因此有了下面的rune的切片,这样返回的将是一个字符而不是一个字节
   
	chars := []rune(text1) //先把字符串转为rune切片
	for _, char := range chars {
		fmt.Println(string(char)) // 在
		fmt.Println(char) //22312
	}
}
