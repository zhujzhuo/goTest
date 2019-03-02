package main

import(
 "fmt"
 "time"
)

var t time.Time
func main() {
     //fmt.Println(1*time.Second - time.Duration(time.Time.Nanosecond())*time.Nanosecond)
     fmt.Println(1*time.Second - time.Duration(t.Nanosecond())*time.Nanosecond)
     second := time.Second
     fmt.Println(int64(second/time.Millisecond))
     fmt.Println(t.Nanosecond())
     fmt.Println(time.Time{})
     fmt.Println(time.Now())
}


