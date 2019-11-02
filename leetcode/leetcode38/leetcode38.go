/*
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
6.     312211
7.     13112221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
注意：整数顺序将表示为一个字符串。



示例 1:
输入: 1
输出: "1"

示例 2:
输入: 4
输出: "1211"
*/

package main

import (
	"fmt"
	//    "strconv"
)

/*
给出一个初始字符串"1"，根据以下规则生成字符串。如果一个字符s[i]与s[i-1]和s[i+1]不同，那么就生成"1"+s[i]的字符串；如果连续n个字符都相同，那么就生成"n"+s[i]的字符串。例如字符串"1"会生成"11"，而"11"会生成"21"，"21"又会生成"1211"，以此类推。
*/
func countAndSay(n int) string {
	//inti,err := strconv.Atoi(str)
	//strs,err := strconv.Itia(int)

	//这里使用byte数组和string容易转化
	if n == 1 {
		return "1"
	}
	bs := []byte{'1'}
	for i := 2; i <= n; i++ {
                fmt.Println(bs)
		bs = say(bs)
	}
	return string(bs)
}
func say(bs []byte) []byte {
	res := make([]byte, 0)
	var x, y int = 0, 1 //当前和下一个index
	//1 11 21 1211 111221
	for x < len(bs) {
		for y < len(bs) && bs[x] == bs[y] {
			y++
		}
		res = append(res, byte(y-x+'0'), bs[x])
		x = y
                fmt.Println(y)
	}
	return res
}

/*
func say(bs []byte) []byte {
    result := make([]byte, 0)
    x, y := 0, 1
    //1 11 21 1211
    for x < len(bs) {
        //取出字节数组中的每一个
        //判断相邻位置的是否是一样
        //当是一样的话，那就继续，找到有多少个是一样的
        //当不是一样的话，那就是一个几
        //这里要保证y不能超过bs的长度，不然会panic

        //相当于有两个index 分别指向当前元素和当前元素的下一个，下一个index会根据比对进行偏移，用来计算重复数字的个数
        for y < len(bs) && bs[x] == bs[y] {
            y++
        }
        //第二个参数是指有多少个一样的
        // (这里需要注意一点，一定要加上'0'，不然字节不对,'0'代表的字节是48，
        // 如果不加上'0'，byte(y-x)就是byte(1),这是不对的)
        //第三个参数是指说出来的那个数
        result = append(result, byte(y-x+'0'), bs[x])

        //index 偏转跳过已经扫描过的元素
        x = y
        y++   // 这里y不自增也可以跑出来结果，不知道为什么? 此处y不自增，在上面内部y的for循环时候，第一次的比较肯定是成立的（因为此时x和y相等），因此y还是会++，也就是说此处没有y++，在内部循环的时候y还是会++，表示的是重复数据的个数
    }
    return result
}
*/
/*
   //byte 测试
   bs := []byte{'1'}
   result := make([]byte, 0)
   result = append(result, byte(1+'0'), bs[0])
   fmt.Println(string(result)) //11
   fmt.Println(byte(1))   //1   这个不是表示数字1，数字1需要如下的表示
   fmt.Println(byte(1+'0')) //49
   fmt.Println(string(byte(1)))   //一个空的输出
   fmt.Println(string(byte(1+'0'))) //1
*/

func main() {
//	fmt.Println(countAndSay(1))
//	fmt.Println(countAndSay(2))
//	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
//	fmt.Println(countAndSay(5))
//	fmt.Println(countAndSay(6))
//	fmt.Println(countAndSay(7))
//	fmt.Println(countAndSay(8))
}
