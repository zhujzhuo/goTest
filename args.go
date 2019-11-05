package main

//import "math"
import "fmt"

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

type Player struct {
	Name  string "name"
	Level int    "level"
	Exp   int    "exp"
	Room  int    "room"
}

func myfunc2(args []int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

// 如果使用.(type)查询类型的变量不是interface{}类型，则在编译时会报错
// cannot type switch on non-interface value a (type string)
// 如果在switch以外地方使用.(type)，则在编译时会提示如下错误：
// use of .(type) outside type switch
// 使用type进行类型查询时，只能在switch中使用，且使用类型查询的变量类型必须是interface{}
// 使用interface{}仍然是类型安全的,
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main() {
	/*
	   两个不同类型的整型数不能直接比
	   var i int32
	   var j int64
	   i,j = 1,2
	   if i==j { // 编译报错
	   }
	*/
	/*
	   字符的内容可以用类似于数组下标的方式获取，但与数组不同，字符的内容不能在初始化后被修改，比如以下的例子
	   str := "Hello world"
	   str[0] = 'X' //   编译会报错
	*/
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	var v5 string = "test"
	MyPrintf(v1, v2, v3, v4)
	i := []int{1, 2, 3, 4, 5}
	myfunc(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("=============== print i:")
	for _, v := range i {
		fmt.Println(v)
	}
	fmt.Println("=============== print i:")
	myfunc2(i)
	MyPrintf(i)
	MyPrintf(v5)
}
