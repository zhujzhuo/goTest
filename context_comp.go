package main
//理解 golang 中的 context（上下文） 包

import "fmt"

/*
//function to print hello
func printHello() {
        fmt.Println("Hello from printHello")
}
func main() {
        //inline goroutine. Define a function inline and then call it.
        go func(){fmt.Println("Hello inline")}()
        //call a function as goroutine
        go printHello()
        fmt.Println("Hello from main")
//如果您运行上面的程序，您只能看到 main 中打印的 Hello,其他两个goroutine可能不能打印完全。 因为它启动了两个 goroutine 并在它们完成前退出了。为了让 main 等待这些 goroutine 执行完，您需要一些方法让这些 goroutine 告诉 main 它们执行完了，那就需要用到通道。
}

*/
//prints to stdout and puts an int on channel
func printHello(ch chan int) {
    fmt.Println("Hello from printHello")
    //send a value on channel
    ch <- 2
}
func main() {
    //make a channel. You need to use the make function to create channels.
    //channels can also be buffered where you can specify size. eg: ch := make(chan int, 2)
    //that is out of the scope of this post.
    ch := make(chan int,2)
    //inline goroutine. Define a function and then call it.
    //write on a channel when done
    go func() {
        fmt.Println("Hello inline")
        //send a value on channel
        ch <- 1
    }()
    //call a function as goroutine
    go printHello(ch)
    fmt.Println("Hello from main")
    //get first value from channel.
    //and assign to a variable to use this value later
    //here that is to print it
    i := <-ch
    fmt.Println("Recieved ", i)
    //get the second value from channel
    //do not assign it to a variable because we dont want to use that
    <-ch
}

//在 Go 语言中 context 包允许您传递一个 "context" 到您的程序。 Context 如超时或截止日期（deadline）或通道，来指示停止运行和返回。例如，如果您正在执行一个 web 请求或运行一个系统命令，定义一个超时对生产级系统通常是个好主意。因为，如果您依赖的API运行缓慢，你不希望在系统上备份（back up）请求，因为它可能最终会增加负载并降低所有请求的执行效率。导致级联效应。这是超时或截止日期 context 派上用场的地方。


