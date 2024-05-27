package main

import "fmt"

/*
// 生成2, 3, 4, ... 到 channel 'ch'中.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i	// Send 'i' to channel 'ch'.
	}
}
// 从管道复制值 'in' 到 channel 'out',
// 移除可整除的数 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in	// 接收值 'in'.
		if i%prime != 0 {
			out <- i	// 传入 'i' 到 'out'.
		}
	}
}
func main() {
	ch := make(chan int)	// Create a newchannel.
	go Generate(ch)	// Launch Generate goroutine.
	for i := 0; i < 600; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
*/

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	num := 0
	for j := 1; j < 100; j++ {
		if isPrime(j) {
			num++
			fmt.Println(j, isPrime(j))
		}
	}
	fmt.Println(num, "个质数")
}
