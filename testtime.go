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
     fmt.Println(time.Now().Year())
     fmt.Println(time.Now().Hour())
     fmt.Println(time.Now().Minute())
     fmt.Println(time.Now().Second())

/*
1s
1000
0
0001-01-01 00:00:00 +0000 UTC
2019-05-17 15:46:32.47081 +0800 CST m=+0.000323119
2019
15
46
32
*/
}


