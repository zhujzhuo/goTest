package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	//a := "hello"
	var a string = "hello"
	fmt.Println("Hello World!")
	fmt.Println(a)
	var str1 string = "hello"
	var str2 string = "123"
	fmt.Println(str1 + str2)
	fmt.Println(len(str1))

	fmt.Println(runtime.GOARCH)
	fmt.Println(strings.ToUpper(runtime.GOOS))

	fmt.Println(os.Getenv("ETCDCOV_ARGS"))
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Getenv("LANG"))

	val := "test,for,a,te if not,print"
	//val := "if not print"
	v := strings.Split(val, ",")
	for _, i := range v {
		fmt.Println(i)
	}
}
