package main

import (
	"fmt"
	"reflect"
)

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

//字符串反转测试
func main() {
	var str1 string = "abcdefg"
	var str2 string = "{}[]"

	fmt.Println(reverseString(str1))
	fmt.Println(str1)
	fmt.Println(reverseString(str2))
	fmt.Println(str2)
	x := "[]"
	t := x[0]
	fmt.Println("type:", reflect.TypeOf(t))
	fmt.Println(string(x[0]))
	fmt.Println(string(x))
}

/*
//map 的赋值和遍历
func main() {
        var mapstring map[string] int
        mapstring = make(map[string]int,20)
        mapstring["name1"] = 1
        mapstring["name2"] = 2
        mapstring["name3"] = 3
	//var mapstring = map[string]int{"name1": 1, "name2": 2, "name3": 3}
	//mapstring := map[string]int{"name1": 1, "name2": 2, "name3": 3}
	for key, value := range mapstring {
		fmt.Printf("%s:%d\n", key, value)
	}
}
*/
