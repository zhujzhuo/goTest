package main

import (
	"fmt"
	"runtime"
	"time"
)
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
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	var v5 string = "test"
	MyPrintf(v1, v2, v3,  v4,v5)
        
        fmt.Println(runtime.GOOS) 
	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Printf("When's Saturday?%s",time.Saturday)
	fmt.Println()
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	//没有条件的switch 可以用来较清晰的编写较长的if-then-else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println("counting")
	defer fmt.Println("done1")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println("done2")
	defer fmt.Println("done3")
	defer fmt.Println("done4")

}
