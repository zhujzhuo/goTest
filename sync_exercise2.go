package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//多个线程同时运行，获得Mutex锁者线程优先执行，其余线程阻塞等待
func testMutex() {
	mutex := sync.Mutex{}
	for i := 0; i < 5; i++ {
		go func(idx int) {
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Println("idx :=", idx)
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Func finish.")
}

//写请求在读锁和写锁时都必须阻塞等待，读请求只在写锁时阻塞等待
func testRWMutex() {
	rwMutex := sync.RWMutex{}
	for i := 0; i < 5; i++ {
		go func(idx int) {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			fmt.Println("Read Mutex :", idx)
		}(i)

		go func(idx int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			fmt.Println("Write Mutex :", idx)
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Func finish.")
}

//Cond：条件变量，等同于linux下的pthread_cond_t
func testCond() {
	cond := sync.NewCond(&sync.Mutex{})

	cond.L.Lock() //①上锁
	defer cond.L.Unlock()

	go func() {
		fmt.Println("go wait lock.")
		cond.L.Lock() //②等Wait解锁

		defer cond.L.Unlock() //⑤解锁后触发Wait
		defer fmt.Println("go unlock.")

		fmt.Println("go locked.")
		cond.Signal() //④触发Wait等待解锁
	}()

	time.Sleep(time.Second)

	fmt.Println("start wait.")
	for {
		cond.Wait() //③可以理解为立刻解锁并触发一个阻塞线程（如果没有阻塞线程则不触发）后立刻再上锁等待Signal信号
		fmt.Println("wait finish.")
		break
	}

	time.Sleep(time.Second)
	fmt.Println("Func finish.")
}

//WaitGroup：组等待
//Add 增加等待计数;Done减少等待计数；当计数为0时触发Wait;
func testWaitGroup() {
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)

		go func(idx int) {
			time.Sleep(time.Second)
			fmt.Println("go : ", idx)
			waitGroup.Done()
		}(i)
	}

	for {
		fmt.Println("start wait.")
		waitGroup.Wait()
		fmt.Println("wait finish.")
		break
	}

	time.Sleep(time.Second)
	fmt.Println("Func finish.")
}

//once 只执行一次以后不再触发
func testOnce() {
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		go func(idx int) {
			once.Do(func() {
				fmt.Println("Do once : ", idx) //这里只执行一次
			})

			fmt.Println("go : ", idx)
		}(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Func finish.")
}

//Map：线程安全map
func testMap() {
	syncMap := sync.Map{}
	for i := 0; i < 10; i++ {
		go func(idx int) {
			//如果没有则保存起来
			_, ok := syncMap.LoadOrStore(idx, " StrVal = "+strconv.FormatInt(int64(idx), 10))
			if !ok {
				fmt.Println("Store idx = ", idx)
			}
		}(i)

		go func(idx int) {
			val, ok := syncMap.Load(idx)
			if ok {
				fmt.Println("Load success idx = ", idx, val)
			} else {
				fmt.Println("Load fail idx = ", idx)
			}
		}(i)

	}

	time.Sleep(5 * time.Second)
	fmt.Println("Func finish.")
}

//Pool：线程安全对象池
func testPool() {
	p := &sync.Pool{
		New: func() interface{} {
			return -1
		},
	}

	for i := 0; i < 10; i++ {
		go func(idx int) {
			p.Put(idx)
		}(i)
	}

	//取出来的对象是无序的
	for i := 0; i < 20; i++ {
		go func() {
			val := p.Get()
			fmt.Println("Get val = ", val)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Func finish.")
}

func main() {
	fmt.Println("start main.")
	testMutex()
	testRWMutex()
	testWaitGroup()
}
