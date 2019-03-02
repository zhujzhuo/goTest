package main

import (
	"fmt"
	"time"
)

func add(x, y int) {
	z := x + y
	fmt.Println("add 内 i的值是:", x)
	fmt.Println(z)
}
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i的值是:", i)
		go add(i, i)
		//协程可以使得for并发执行，如果这个函数有返回值，那么这个返回值会被丢弃
		//如下的sleep可以看到顺序执行的输出,当然在go中要看到协程输出需要引入goroutine的并发通信
		//
		time.Sleep(time.Duration(1)*time.Second)
	}

	const (
		c0 = iota //0
		c1        //1
		c2        //2
	)
	const (
		a0 = 1 << iota //1
		a1             //2
		a2             //4
	)

	fmt.Println("c0=", c0) //0
	fmt.Println("c1=", c1) //1
	fmt.Println("c2=", c2) //2
	fmt.Println("a0=", a0) //1
	fmt.Println("a1=", a1) //2
	fmt.Println("a2=", a2) //4

}
