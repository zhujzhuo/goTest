package main
//Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct
//使用传值或者传地址都行，会自动的识别。
import "fmt"

type rect struct {
    width, height int
}

//传地址
func (r *rect) area() int {
    return r.width * r.height
}
//传值
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

//如上的两种都可以使用，有些喜欢使用值，有些喜欢使用地址定义，可以参见interface2.go的代码，直接使用了值来定义其interface

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
/*
area:  50
perim: 30
area:  50
perim: 30
*/
