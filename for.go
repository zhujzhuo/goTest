package main

import "fmt"
import "math"
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	sum := 0
	for i := 0; i < 101; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	//可以省略1  3 条件
	for ; sum2 < 1000; {
		sum2 += sum2
	}
	fmt.Println(sum2)

        //C 的 while 在 Go 中叫做 `for`
        fmt.Println("===========")
        sum3 :=1
        for sum3 < 1000 {
	        fmt.Println(sum3)
		sum3 += sum3
	}
	fmt.Println(sum3)
	fmt.Println(sqrt(2), sqrt(-4))
	
}
