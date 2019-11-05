package main

import (
	"fmt"
	"sync"
)

/*
sync 包提供了互斥锁这类的基本的同步原语.除 Once 和 WaitGroup 之外的类型大多用于底层库的例程。更高级的同步操作通过信道与通信进行。
type Cond
    func NewCond(l Locker) *Cond
    func (c *Cond) Broadcast()
    func (c *Cond) Signal()
    func (c *Cond) Wait()
type Locker
type Mutex
    func (m *Mutex) Lock()
    func (m *Mutex) Unlock()
type Once
    func (o *Once) Do(f func())
type Pool
    func (p *Pool) Get() interface{}
    func (p *Pool) Put(x interface{})
type RWMutex
    func (rw *RWMutex) Lock()
    func (rw *RWMutex) RLock()
    func (rw *RWMutex) RLocker() Locker
    func (rw *RWMutex) RUnlock()
    func (rw *RWMutex) Unlock()
type WaitGroup
    func (wg *WaitGroup) Add(delta int)
    func (wg *WaitGroup) Done()
    func (wg *WaitGroup) Wait()
而golang中的同步是通过sync.WaitGroup来实现的．WaitGroup的功能：它实现了一个类似队列的结构，可以一直向队列中添加任务，当任务完成后便从队列中删除，如果队列中的任务没有完全完成，可以通过Wait()函数来出发阻塞，防止程序继续进行，直到所有的队列任务都完成为止．
WaitGroup总共有三个方法：Add(delta int),　Done(),　Wait()。
Add:添加或者减少等待goroutine的数量
Done:相当于Add(-1)
Wait:执行阻塞，直到所有的WaitGroup数量变成0
*/

var waitgroup sync.WaitGroup

func Afunction(shownum int) {
	fmt.Println(shownum)
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}

func main() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go Afunction(i)
	}
	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
}
