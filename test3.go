package main

import (
"fmt"
)

func add(x,y int) int{
    return x+y
}

func main(){
    var x,y int=3,4
	z := add(x,y)
	fmt.Println(z)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	
	// int  
	var a []int
	//var a [5]int
    a = append(a, 0)

    // double 
	var double = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, v := range double {
		fmt.Printf("2*%d = %d\n", v, v*2)
	}
    // 数组
	var myArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    //myArray2 := [...] int {1,2,3,4,5}
    //var myArray3 = [...] int {1,2,3}
     
	for _,val := range myArray{
		fmt.Println(val)
	}
    //注意slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符。
	//切片初始化：
	s1 :=[] int {1,2,3 }
	//直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
	s2 :=make([]int,len,cap)
	//通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片

	

}
