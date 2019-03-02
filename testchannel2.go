package main

import "fmt"

func Afuntion(ch chan int) {
	fmt.Println("finish")
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan int) //无缓冲的channel
	go Afuntion(ch)
	ch <- 1
	//只是把上面的两行的代码顺序对调一下
	// 输出结果：
	// 死锁，无结果
        //因此必须go协程在前，解锁协程在后．不能是主进程被hang住
}
