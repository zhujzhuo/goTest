package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
        p1 = &v1
)

func main() {
	fmt.Println(v1, p, v2, v3,p1)  //{1 2} &{1 2} {1 0} {0 0} &{1 2}
        fmt.Printf("p1.X:%v,p1.Y:%v",p1.X,p1.Y)
        fmt.Println()
        fmt.Printf("p.X:%v,p.Y:%v",p.X,p.Y)
        fmt.Println()
}


