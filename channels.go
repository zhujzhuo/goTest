package main

import (
	"fmt"
	//  "time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将sum送入 c
}
func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int, 2)
	go sum(a[:len(a)/2], c)
	// 这里加上time sleep保证x y的复制顺序，否则go多线程执行。当然加上sleep之后就失去了并发执行的优势了
	// time.Sleep(time.Duration(1)*time.Second)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // 从 c 中获取
	fmt.Println(x, y, x+y)

	//channel 可以是带缓冲的。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：
	//ch := make(chan int, 100)
	//向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞。
	//修改例子使得缓冲区被填满，然后看看会发生什么。
	d := make(chan int, 2)
	d <- 1
	fmt.Println(<-d)
	d <- 2
	d <- 3
	fmt.Println(<-d)
	fmt.Println(<-d)
	d <- 4
	d <- 5
	fmt.Println(<-d)
	fmt.Println(<-d)
}

//channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。
//ch <- v    // 将 v 送入 channel ch。
//v := <-ch  // 从 ch 接收，并且赋值给 v。（“箭头”就是数据流的方向。）
//和 map 与 slice 一样，channel 使用前必须创建：
//ch := make(chan int)
//默认情况下，在另一端准备好之前，发送和接收都会阻塞。这使得 goroutine
//可以在没有明确的锁或竞态变量的情况下进行同步。

//从带缓冲的channel中读取数据可以使用与常规非缓冲channel完全一致的方法，但我们也可以使用range关键来实现更为简便的循环读取:
//for i := range c {
//   fmt.Println("Received:", i)
//}

// 超时处理
//timeout := make(chan bool, 1)
//go func() {
//   time.Sleep(1e9) // 等待1秒钟
//   timeout <- true
//}()
// 然后我们把timeout这个channel利用起来
//select {
//   case <-ch:
//   // 从ch中读取到数据
//   case <-timeout:
// 一直没有从ch中读取到数据，但从timeout中读取到了数据
// 这样使用select机制可以避免永久等待的问题，因为程序会在timeout中获取到一个数据后继续执行，无论对ch的读取是否还处于等待状态，从而达成1s超时的效果
//}

//var ch1 chan int // ch1是一个正常的channel，不是单向的
//var ch2 chan<- float64  // ch2是单向channel，只用于写float64数据
//var ch3 <-chan int // ch3是单向channel，只用于读取int数据

//ch4 := make(chan int)
//ch5 := <-chan int(ch4) // ch5就是一个单向的读取channel
//ch6 := chan<- int(ch4) // ch6 是一个单向的写入channel
