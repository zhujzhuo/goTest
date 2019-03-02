/*
package main 
import "fmt"

type gameObject struct {
    name string
}
func (o *gameObject) Name() string{
    return o.name
}
func (o *gameObject) Attack() {
    fmt.Printf("GameObject %s do not know how to attack\n", o.name)
}

type player struct {
    gameObject
}
func (p *player) Attack(){
    fmt.Printf("player %s attacks\n", p.name)
}

type monster struct {
    gameObject
}

type challenger interface {
    Name() string
    Attack()
}

func main(){
    p := &player{gameObject : gameObject{name:"devgl"}}
    m := &monster{gameObject : gameObject{name:"slime"}}
    Attack(p)
    Attack(m)
}

func Attack(c challenger) {
    c.Attack()
}


*/
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a,b Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	b = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v

	fmt.Println(a.Abs())
	fmt.Println(b)
       
        var phone Phone
        phone = new(NokiaPhone)
        phone.call()
        phone = new(IPhone)
        phone.call()
}

