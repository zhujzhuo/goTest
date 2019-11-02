package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	a := "hello"
	fmt.Println("Hello World!")
	fmt.Println(a)
	var str1 string = "hello"
	var str2 string = "123"
	fmt.Println(str1 + str2)
	fmt.Println(len(str1))
	fmt.Println(runtime.GOARCH)
	fmt.Println(os.Getenv("ETCDCOV_ARGS"))
}
