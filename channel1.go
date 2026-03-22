package main

import (
	"fmt"
	"time"
)

func Afunctionin(ch chan string) {
	ch <- "test channel for one example"
	fmt.Println("finishin")
}
func Afunctionout(ch chan int) {
	fmt.Println("finishout")
	<-ch
}
func main() {

	ch2 := make(chan int) //无缓冲的channel
	go Afunctionout(ch2)  //先使用协程去读取channel
	// sleepDuration := time.Duration(rand.IntN(5)+1) * time.Second
	// time.Sleep(sleepDuration)
	ch2 <- 1            //再写入数据
	time.Sleep(1e9 * 4) // 等待1秒钟
	// 输出结果：
	// finishout   之后会sleep几秒钟
	ch := make(chan string) //无缓冲的channel
	go Afunctionin(ch)      //先使用协程去写channel
	fmt.Println(<-ch)       //再读数据
	// 输出结果：
	// finishin
	// test channel for one example

}
