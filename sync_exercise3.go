package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	sum := 0
	sum2 := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		sum += i
		go fmt.Println("sum:", sum)
		wg.Done()
	}
	for {
		fmt.Println("start wait.")
		wg.Wait()
		fmt.Println("wait finish.")
		break
	}
	rwMutex := sync.RWMutex{}
	for j := 0; j < 5; j++ {
		go func(idx int) { //使用了go使得后面的j不确定
			rwMutex.Lock()
			defer rwMutex.Unlock()
			sum += j //这个地方的j是不确定的，可能是0 1 2 3 4 5 中的任何一个
			fmt.Println("Write Mutex :", idx)
			fmt.Println("sum&j", sum, j)
		}(j)
		func(idx int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			sum2 += j //这个地方的j是确定的
			fmt.Println("Write Mutex :", idx)
			fmt.Println("sum2&j", sum2, j)
		}(j)

	}
	// rwMutex.Lock()
	// defer rwMutex.Unlock()
	// fmt.Println("Write Mutex end:")
	time.Sleep(2 * time.Second)
	fmt.Println("Func finish.")
	fmt.Println("sum:", sum)
	fmt.Println("sum2:", sum2)
}
