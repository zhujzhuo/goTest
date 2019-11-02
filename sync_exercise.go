package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
func (o *Once) Do(f func())

func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()

*/
func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	y := 0
	for i := 0; i < 10; i++ {
		//time.Sleep(time.Duration(3)*time.Second)
		//fmt.Println(len(done))
		//fmt.Println(cap(done))
		go func() {
			once.Do(onceBody) //这个只执行一次
			done <- true      //这个赋值会在后面fmt.Println(<-done) 解锁之后再写入
			y++
		}()
	}
	fmt.Println(y) // 0
	for i := 0; i < 10; i++ {
		fmt.Println(y)      // 这个地方打印值无法预估,或前或后于chan的赋值
		fmt.Println(<-done) //最终打印10个true
	}
	fmt.Println(y) // 10

	//wait group 用来等待一组goroutines的结束，在主Goroutine里声明，并且设置要等待的goroutine的个数，每个goroutine执行完成之后调用 Done，最后在主Goroutines 里Wait即可
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			fmt.Println(http.Get(url))
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()

	var mutex sync.Mutex
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")
	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not lock: ", i)
			mutex.Lock()
			fmt.Println("Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock: ", i)
			mutex.Unlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		<-c
	}
	/*
	   Lock the lock
	   The lock is locked
	   Not lock:  3
	   Not lock:  0
	   Not lock:  1
	   Not lock:  2
	   Unlock the lock
	   Locked:  3
	   Unlock the lock:  3
	   Locked:  0
	   Unlock the lock:  0
	   Locked:  1
	   Unlock the lock:  1
	   Locked:  2
	   Unlock the lock:  2
	*/
}
