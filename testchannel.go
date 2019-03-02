package main
import  (
"fmt"
)

func main(){
c := make(chan int,20)
defer  close(c)
sum :=0
for i:=0;i<10;i++{
    c <-i
    sum ++
}

for i:=0;i<10;i++{
    fmt.Println(<-c)
}
fmt.Println("The chan's length is:",sum)
fmt.Println("The chan's cap is:",cap(c))
}

