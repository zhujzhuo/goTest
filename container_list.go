package main
import (
    "fmt"
    "container/list"
)
func main() {
    // 生成队列
    l := list.New()
    // 入队, 压栈
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    l.PushBack(4)

    fmt.Println(l.Len())
    // 出队
    i1 := l.Front()
    l.Remove(i1)
    fmt.Printf("%d\n", i1.Value)
    i2 := l.Front()
    l.Remove(i2)
    fmt.Printf("%d\n", i2.Value)
    // 出栈
    i4 := l.Back()
    l.Remove(i4)
    fmt.Printf("%d\n", i4.Value)
    i5 := l.Back()
    l.Remove(i5)
    fmt.Printf("%d\n", i5.Value)
}
