package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5]
	fmt.Println("E of myArray:")
	for i, el := range myArray {
		fmt.Print(i, ":", el, " ")
	}
	fmt.Println()
	fmt.Println("E of mySlice:")
	for _, el := range mySlice {
		fmt.Print(el, " ")
	}
	fmt.Println()
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	for i, el := range slice1 {
		fmt.Print(i, ":", el, "  ")
	}
	fmt.Println()
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println("======")
	for _, el := range slice2 {
		fmt.Print(el, " ")
	}
	//创建一个slice
	yourSlice := make([]int, 5, 10)
	fmt.Println()
	fmt.Println("len(yourSlice):", len(yourSlice))
	fmt.Println("cap(yourSlice):", cap(yourSlice))

	mySlice2 := []int{8, 9, 10}
	yourSlice = append(slice1, mySlice2...)
	//需要注意的是,我们在第二个参数mySlice2后面加了三个点,即一个省略号,如果没有这个省略号的话,会有编译错误,
	//因为按append()的语义,从第二个参数起的所有参数都是待附加的元素。因为yourSlice中的元素类型为int,
	//所以直接传递mySlice2是行不通的。加上省略号相当于把mySlice2包含的所有元素打散后传入。
	for i, el := range yourSlice {
		fmt.Print(i, ":", el, "  ")
	}
	fmt.Println()
	//s[lo:hi] 表示从 lo 到 hi-1 的 slice 元素，含两端。 从0开始
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4]) //显示 3 5 7
	fmt.Println("p[:1] ==", p[:1])   //显示 2

	// 省略下标代表从 0 开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("p[4:] ==", p[4:])

	var z []int
	c := make([]int, 5, 10)
	fmt.Println(z, len(z), cap(z))
	fmt.Println(c, len(c), cap(c))
	if z == nil {
		fmt.Println("nil!")
	}
}
