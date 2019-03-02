package main

import "fmt"

func main(){
    str := "helloWorld在"
    fmt.Println(len(str)) //10+3  
    fmt.Println(str[0]) //104
    fmt.Printf("%c,%c,%c\n",str[1],str[3],str[7]) //e l r
    fmt.Println(1%10)
    fmt.Println(1/10)
    var strs []string = []string{""}
    fmt.Println(len(strs))
    fmt.Println(strs[:1])
    fmt.Println(strs[:0])
    fmt.Println(str[:1])
}
