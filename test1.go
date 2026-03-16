package main

import (
	"fmt"
	"math"
	"time"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4)) //5

	fmt.Println(compute(hypot))    // 5
	fmt.Println(compute(math.Pow)) //math.Pow(x, y), x的y次方   3*3*3*3=81

	var a int = 1234
	fmt.Println("Hello World!")
	fmt.Println(a)
	var str1 string = "hello"
	var str2 string = "123"
	fmt.Println(time.Now())
	time.Sleep(time.Duration(1000) * time.Millisecond) //2s
	fmt.Println(time.Now())
	fmt.Println(str1 + str2)
	fmt.Println(len(str1))

	var tablename string = "tablename_test"
	var myArray []string = []string{"test1", "test2"}
	myArray = append(myArray, tablename)
	fmt.Println(myArray)

	customTime := time.Now().Format("15:04:05")
	fmt.Println(customTime)

}
