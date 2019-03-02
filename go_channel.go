package main
import "fmt"


func main (){
var ch = make(chan int,20)
for i := 0; i < 10; i++ {
	ch <- i
}
close (ch)
for i:=range ch {
   fmt.Println(i)
}
}
