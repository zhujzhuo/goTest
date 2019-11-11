/*
数组是内置(build-in)类型,是一组同类型数据的集合，它是值类型，通过从0开始的下标索引访问元素值。在初始化后长度是固定的，无法修改其长度。当作为方法的参数传入时将复制一份数组而不是引用同一指针。数组的长度也是其类型的一部分，通过内置函数len(array)获取其长度。
注意：和C中的数组相比，又是有一些不同的

1. Go中的数组是值类型，换句话说，如果你将一个数组赋值给另外一个数组，那么，实际上就是将整个数组拷贝一份
2. 如果Go中的数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针。这个和C要区分开。因此，在Go中如果将数组作为函数的参数传递的话，那效率就肯定没有传递指针高了。
3. array的长度也是Type的一部分，这样就说明[10]int和[20]int是不一样的
数组初始化：
var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
myArray2 := [...] int {1,2,3,4,5}
var myArray3 = [...] int {1,2,3}

切片中有两个概念：一是len长度，二是cap容量，长度是指已经被赋过值的最大下标+1，可通过内置函数len()获得。容量是指切片目前可容纳的最多元素个数，可通过内置函数cap()获得。切片是引用类型，因此在当传递切片时将引用同一指针，修改值将会影响其他的对象。
slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，例如上面的aSlice和bSlice，如果修改了aSlice中元素的值，那么bSlice相对应的值也会改变
aSlice = array[3:7] // aSlice包含元素: d,e,f,g，len=4，cap=7
bSlice = aSlice[1:3]// bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
bSlice = aSlice[:3] // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有:

切片可以通过数组来初始化，也可以通过内置函数make()初始化 .初始化时len=cap,在追加元素时如果容量cap不足时将按len的2倍扩容

切片初始化：
s :=[] int {1,2,3 }
直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

s :=make([]int,len,cap)
通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片

注意slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符。

// 声明一个含有10个元素元素类型为byte的数组
var ar = [10]byte{'a','b','c','d','e','f','g','h','i','j'}
// 声明两个含有byte的slice
var a, b []byte
// a指向数组的第3个元素开始，并到第五个元素结束，
a = ar[2:5]
//现在a含有的元素: ar[2]、ar[3]和ar[4]

// b是数组ar的另一个slice
b = ar[3:5]
// b的元素是：ar[3]和ar[4]

*/
package main

import "fmt"

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5]

	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice: ")
	for i, v := range mySlice {
		fmt.Println("mySlice[", i, "] =", v)
	}
	fmt.Println()
	mySlice2 := []int{11, 12, 13}
	//第二个参数mySlice2后面加了三个点，如果没有这个的话，会有编译错误，因为 append()的语义，从第二个参数起的所有参数都是
	//待附加的元素。因为mySlice中的元素类型为int，所以直接传递mySlice2 type []int 是行不通的。加上省略号相当于把mySlice2包
	//含的所有元素打散后传入
	mySlice = append(mySlice, mySlice2...)
	mySlice = append(mySlice, 14, 15, 16)
	for i, v := range mySlice {
		fmt.Println("mySlice[", i, "] =", v)
	}

	//mySlice1 := make([]int, 5)
	//与数组相比，数组切片多了一个存储能力(capacity)的概念， 元素个数和cap之间可以是两个不同的值
	//mySlice2 := make([]int, 5, 10)
	//mySlice3 := []int{1, 2, 3, 4, 5}
	fmt.Println("数组和切片的使用:")
	array := [...]float64{7.0, 8.5, 9.1, 1.1, 2.2, 3.3}
	aslice := array[2:5]
	bslice := aslice

	bslice[1] = 5.5
	for _, i := range aslice {
		fmt.Println(i)
	}

	for _, i := range bslice {
		fmt.Println(i)
	}
	fmt.Println("=====================")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a)
	//二维数组
	type LinesOfText [][]byte
	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	fmt.Println(text)
	type LinesOfInt [][]int
	intvar := LinesOfInt{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(intvar)
	//数组切片的初始化，必须在使用之前做make的初始化
	// 分配顶层切片。
	//picture := make([][]uint8, YSize) // 每 y 个单元一行。
	// 遍历行，为每一行都分配切片
	//for i := range picture {
	//    picture[i] = make([]uint8, XSize)
	//}
}
