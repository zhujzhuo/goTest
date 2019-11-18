package main

import (
    "fmt"
    "time"
)

func main() {

    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)
    //2019-11-18 11:28:16.599939 +0800 CST m=+0.000163882
    //Note that there is no UnixMillis, so to get the milliseconds since epoch you’ll need to manually divide from nanoseconds.
    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)
    //You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0,nanos))
    //2019-11-18 11:28:16 +0800 CST
    //2019-11-18 11:28:16.599939 +0800 CST    
}
