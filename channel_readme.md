Channel机制：

相对sync.WaitGroup而言，golang中利用channel实现同步则简单的多．channel自身可以实现阻塞，其通过<-进行数据传递，channel是golang中一种内置基本类型，对于channel操作只有４种方式：

创建channel(通过make()函数实现，包括无缓存channel和有缓存channel);
向channel中添加数据（channel<-data）;
从channel中读取数据（data<-channel）;
关闭channel(通过close()函数实现，关闭之后无法再向channel中存数据，但是可以继续从channel中读取数据）

channel分为有缓冲channel和无缓冲channel,两种channel的创建方法如下:
var ch = make(chan int) //无缓冲channel,等同于make(chan int ,0)
var ch = make(chan int,10) //有缓冲channel,缓冲大小是10

其中无缓冲channel在读和写是都会阻塞，而有缓冲channel在向channel中存入数据没有达到channel缓存总数时，可以一直向里面存，直到缓存已满才阻塞．由于阻塞的存在，所以使用channel时特别注意使用方法，防止死锁的产生．例子如下：

无缓存channel:
package main
import "fmt"
func Afunction(ch chan int) {
	fmt.Println("finish")
	<-ch
}
func main() {
	ch := make(chan int) //无缓冲的channel
	go Afunction(ch)
	ch <- 1
	// 输出结果：
	// finish
}
代码分析：首先创建一个无缓冲channel　ch,　然后执行　go Afunction(ch),此时执行＜-ch，则Afunction这个函数便会阻塞，不再继续往下执行，直到主进程中ch<-1向channel　ch 中注入数据才解除Afunction该协程的阻塞．
package main

import "fmt"

func Afunction(ch chan int) {
	fmt.Println("finish")
	<-ch
}

func main() {
	ch := make(chan int) //无缓冲的channel
	//只是把这两行的代码顺序对调一下
	ch <- 1
	go Afunction(ch)

	// 输出结果：
	// 死锁，无结果
}
代码分析：首先创建一个无缓冲的channel,　然后在主协程里面向channel　ch 中通过ch<-1命令写入数据，则此时主协程阻塞，就无法执行下面的go Afunctions(ch),自然也就无法解除主协程的阻塞状态，则系统死锁

总结：
对于无缓存的channel,放入channel和从channel中向外面取数据这两个操作不能放在同一个协程中，防止死锁的发生；同时应该先利用go 开一个协程对channel进行操作，此时阻塞该go 协程，然后再在主协程中进行channel的相反操作（与go 协程对channel进行相反的操作），实现go 协程解锁．即必须go协程在前，解锁协程在后．

带缓存channel:
对于带缓存channel，只要channel中缓存不满，则可以一直向 channel中存入数据，直到缓存已满；同理只要channel中缓存不为０，便可以一直从channel中向外取数据，直到channel缓存变为０才会阻塞．

由此可见，相对于不带缓存channel，带缓存channel不易造成死锁，可以同时在一个goroutine中放心使用，

close():
close主要用来关闭channel通道其用法为close(channel)，并且是在生产者的地方关闭channel，而不是在消费者的地方关闭．并且关闭channel后，便不可再向channel中继续存入数据，但是可以继续从channel中读取数据．例子如下：

package main

import "fmt"

func main() {
    var ch = make(chan int, 20)
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
    //ch <- 11 //panic: runtime error: send on closed channel
    for i := range ch {
        fmt.Println(i)　//输出０　１　２　３　４　５　６　７　８　９
    }
}


channel阻塞超时处理：
goroutine有时候会进入阻塞情况，那么如何避免由于channel阻塞导致整个程序阻塞的发生那？解决方案：通过select设置超时处理，具体程序如下：

package main

 import (
    "fmt"
    "time"
)

func main() {
    c := make(chan int)
    o := make(chan bool)
    go func() {
        for {
            select {
            case i := <-c:
                fmt.Println(i)
            case <-time.After(time.Duration(3) * time.Second):    //设置超时时间为３ｓ，如果channel　3s钟没有响应，一直阻塞，则报告超时，进行超时处理．
                fmt.Println("timeout")
                o <- true
                break
            }
        }
    }()
    <-o
}



golang 并发总结：
并发两种方式：sync.WaitGroup，该方法最大优点是Wait()可以阻塞到队列中的所有任务都执行完才解除阻塞，但是它的缺点是不能够指定并发协程数量．
channel优点：能够利用带缓存的channel指定并发协程goroutine，比较灵活．但是它的缺点是如果使用不当容易造成死锁；并且他还需要自己判定并发goroutine是否执行完．

但是相对而言，channel更加灵活，使用更加方便，同时通过超时处理机制可以很好的避免channel造成的程序死锁，因此利用channel实现程序并发，更加方便，更加易用．
