package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        if sig == syscall.SIGINT{
           fmt.Println("syscall.SIGINT")
        }
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
/*
When we run this program it will block waiting for a signal. By typing ctrl-C (which the terminal shows as ^C) we can send a SIGINT signal, causing the program to print interrupt and then exit.

awaiting signal
^C
interrupt
exiting

*/
