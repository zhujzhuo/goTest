package main

import (
	"fmt"
//      "math"
        "strconv"
        "strings"
        "encoding/hex"
)
/*
func Sum(a *[6]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return sum
}
*/
func main() {
/*
	var v [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	var myArray2 = [...]int{1, 2, 3, 4, 5}

	for _, i := range v {
		fmt.Println(i)
	}

	for _, i := range myArray2 {
		fmt.Println(i)
	}

	array := [...]float64{7.0, 8.5, 9.1, 1.1, 2.2, 3.3}
	x := Sum(&array) // 注意显式的取址操作
	fmt.Println(x)

	aslice := array[2:5]
	bslice := aslice

	bslice[1] = 5.5
	for _, i := range aslice {
		fmt.Println(i)
	}

	for _, i := range bslice {
		fmt.Println(i)
	}

	type ParseOption int

	const (
		Second      ParseOption = 1 << iota // Seconds field, default 0
		Minute                              // Minutes field, default 0
		Hour                                // Hours field, default 0
		Dom                                 // Day of month field, default *
		Month                               // Month field, default *
		Dow                                 // Day of week field, default *
		DowOptional                         // Optional day of week field, default *
		Descriptor                          // Allow descriptors such as @monthly, @weekly, etc.
	)
	fmt.Println(Second)
	fmt.Println(Minute)
	fmt.Println(Hour)
	fmt.Println(Dom)
	fmt.Println(Month)
	fmt.Println(Dow)
	fmt.Println(DowOptional)
	fmt.Println(Descriptor)
        const starBit uint64 = 1 << 63
        fmt.Println(starBit)
        var s string = "abcd"
        fmt.Println(len(s))
        fmt.Println(string(' '))
        
        for i :=10;i>0;i--{
            
            fmt.Printf("%d\n",i)
        }
        fmt.Println(math.Modf(5.26))
        fmt.Println(math.Floor(5.8)) 
        fmt.Println(math.Ceil(5.8))
        fmt.Println(math.Floor(math.Sqrt(8)))
        var aint int = 8 
        fmt.Println(float64(aint))
*/
        var nums1 []int = []int{1,2,3,0,0,0,0}
        var nums2 []int = []int{2,3,4,5}
        m,n := 3,4
        for i,j :=m,0;i<m+n;i++{
           nums1[i]= nums2[j]
           j++
        } 
        fmt.Println(nums1) //[1 2 3 2 3 4 5]
        //排序nums1
        var str string = "12321"
        var count int = 3
        fmt.Println(string(str[1])+"3") //23
        fmt.Println(string(count)) // 没法直接转化输出是空
        fmt.Println(string(count) + string(str[2])) //3
        fmt.Println("3" + string(str[2])) //33
        //字符串和int的转化，需要用到strconv库
        b,_ := strconv.Atoi(str) 
        fmt.Println(b) //12321
        fmt.Println(strconv.Itoa(count)) //3
        fmt.Println("============")
        fmt.Println(str[:0]) //""

        //byte 测试 
        bs := []byte{'1'}
        result := make([]byte, 0) 
        result = append(result, byte(1+'0'), bs[0])
        fmt.Println(string(result)) //11
        fmt.Println(byte(1))   //1   这个不是表示数字1，数字1需要如下的表示
        fmt.Println(byte(1+'0')) //49
        fmt.Println(string(byte(1)))   //一个空的输出
        fmt.Println(string(byte(1+'0'))) //1
        
        //
        var ret [][]int = [][]int{{1},{2,2},{3,4,4,3}}
        var res [][]int
        fmt.Println(len(ret))
        for i:=0;i<len(ret);i++{
            fmt.Println(ret[len(ret)-1-i])
            res = append(res,ret[len(ret)-1-i])
        }
        for i:=0;i<len(ret);i++{
            fmt.Println(res[len(res)-1-i])
        }
        fmt.Println(-324%10)
        //
        var sanitizeQuotesRegexp = "testtestvar"
        fmt.Println(sanitizeQuotesRegexp)    
        //
        var filepath = "/home/mysql/data1"
        fmt.Println(strings.LastIndex("/home/mysql/data1","/"))  // 11
        fmt.Println(filepath[0:12])
  
        src := []byte(" ")  //空格32
        encodedStr := hex.EncodeToString(src)
        fmt.Println(src) //[32]
        fmt.Println(encodedStr) //20
     
        test, _ := hex.DecodeString("\xff\xff") // []
        fmt.Println(test) //[]
}
   
