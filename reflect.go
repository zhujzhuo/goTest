package main

import (
   "fmt"
   "reflect"
)

func main() {
    var y float64 = 3.4
    v := reflect.ValueOf(y)
    fmt.Println("settability of v:" , v.CanSet())
    //v.SetFloat(4.1)
    fmt.Println("type:",reflect.TypeOf(y))
    fmt.Println("value:",reflect.ValueOf(y))
    
    var x float64 = 3.4
    p := reflect.ValueOf(&x) // 注意:得到X的地址 
    fmt.Println("type of p:", p.Type()) 
    fmt.Println("settability of p:" , p.CanSet())
    v1 := p.Elem()
    fmt.Println("settability of v1:" , v1.CanSet())
    fmt.Println("type of v1:" , v1.Type())
    v1.SetFloat(7.1)
    fmt.Println(v1.Interface())
    fmt.Println(x)
    fmt.Println(v1)
    

    type T struct { 
       A int
       B string 
     }
    t := T{203, "mh203"}
    s := reflect.ValueOf(&t).Elem() 
    typeOfT := s.Type()
    fmt.Println("typeOfT:",s.Type())
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d: %s %s = %v\n", i,typeOfT.Field(i).Name, f.Type(), f.Interface())
    }
}
