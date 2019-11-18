package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    //rand.Intn returns a random int n, 0 <= n < 100.
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    //rand.Float64 returns a float64 f, 0.0 <= f < 1.0.
    fmt.Println(rand.Float64())
    
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()
    //给NewSource 多次传一个变化的值，生成的数字才是random的，不然多次都是一样的
    //The default number generator is deterministic, so it’ll produce the same sequence of numbers each time by default. To produce varying sequences, give it a seed that changes. Note that this is not safe to use for random numbers you intend to be secret, use crypto/rand for those.
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()
    
    //If you seed a source with the same number, it produces the same sequence of random numbers.
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
    fmt.Println()
/*
81,87   多次不变
0.6645600532184904
7.1885709359349015,7.123187485356329
34,61   多次变化
5,87
5,87
*/
}
