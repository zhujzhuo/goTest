package main
//Timers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals. Here’s an example of a ticker that ticks periodically until we stop it.

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}
/*
Tick at 2019-11-16 14:56:18.647331 +0800 CST m=+0.500235812
Tick at 2019-11-16 14:56:19.147392 +0800 CST m=+1.000281782
Tick at 2019-11-16 14:56:19.652393 +0800 CST m=+1.505267934
Ticker stopped
*/
