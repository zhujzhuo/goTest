package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        utu := strings.ToUpper(scanner.Text())
        utl := strings.ToLower(scanner.Text())
        fmt.Println(utu)
        fmt.Println(utl)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}

/*
echo 'hello'   > /tmp/lines
echo 'filter' >> /tmp/lines
cat /tmp/lines | go run line-filters.go
HELLO
hello
FILTER
filter
*/
