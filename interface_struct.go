package main

import (
	"fmt"
)

type outstruct struct{
     data  int
     fort  Tests
}

type Tests interface{
     outstr()
}

type Test1 struct {
	out string
        event int
}

type Test2 struct {
	out string
        event int
}

func (t *Test1) outstr() {
	fmt.Println("Test1 out")
}
func (t *Test2) outstr() {
	fmt.Println("Test2 out")
}

func main() {
        t := outstruct{data:1,fort:&Test1{out:"Test1",event:1}}
        fmt.Println("======in Test1")
        t.fort.outstr()
        t2 := outstruct{data:2,fort:&Test2{out:"Test2",event:2}}
        fmt.Println("======in Test2")
        t2.fort.outstr()

        var f,f2  Tests
        
        f = &Test1{out:"Test1",event:3}
        f2 = &Test2{out:"Test2",event:4}


        f3 := f2.(*Test2)
     
        t3 := outstruct{data:3,fort:f}
        fmt.Println("======in Test1")
        t3.fort.outstr()
        t4 := outstruct{data:2,fort:f2}
        fmt.Println("======in Test2")
        t4.fort.outstr()
        t5 := outstruct{data:2,fort:f3}
        fmt.Println("======in Test2")
        t5.fort.outstr()
       
}
