package main

import "fmt"

func Afunctionin(ch chan string) {
	ch <- "test channel for one example"
	fmt.Println("finish")
}
func Afunctionout(ch chan int) {
	fmt.Println("finish")
	<-ch
}
func main() {
	ch := make(chan string) //无缓冲的channel
	go Afunctionin(ch)      //先使用协程去写channel
	fmt.Println(<-ch)       //再读数据
	// 输出结果：
	// finish
	// test channel for one example

	ch2 := make(chan int) //无缓冲的channel
	go Afunctionout(ch2)  //先使用协程去读取channel
	ch2 <- 1              //再写入数据
	// 输出结果：
	// finish
}
