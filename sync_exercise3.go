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
		go func(idx int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			sum += idx
			fmt.Println("Write Mutex :", idx)
			fmt.Println("sum&idx", sum, idx)
		}(j)
		func(idx int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			sum2 += idx
			fmt.Println("Write Mutex :", idx)
			fmt.Println("sum2&idx", sum2, idx)
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
