package main

import (
    "crypto/sha1"
    "fmt"
)

func main() {
    s := "sha1 this string"
    //The pattern for generating a hash is sha1.New(), sha1.Write(bytes), then sha1.Sum([]byte{}). Here we start with a new hash.
    h := sha1.New()

    h.Write([]byte(s))

    bs := h.Sum(nil)

    fmt.Println(s)
    //%x provides hex encoding.
    fmt.Printf("%x\n", bs)
    //sha1 this string
    //cf23df2207d99a74fbe169e3eba035e633b65d94
}
