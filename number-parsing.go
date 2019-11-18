package main

import (
    "fmt"
    "strconv"
    "reflect"
)
/*
func typeof(v interface{}) string {
    //To print the type of a value, use %T
    return fmt.Sprintf("%T", v)
}
*/
func typeof(v interface{}) string {
    return reflect.TypeOf(v).String()
}

func main() {
    //With ParseFloat, this 64 tells how many bits of precision to parse.
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    f2, _ := strconv.ParseFloat("1.234", 1)
    fmt.Println(f2)
    //For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
    //bitSize=64表示转换后的值的类型为int64或uint64
    //base参数表示以什么进制的方式去解析给定的字符串，有效值为0、2-36。当base=0的时候，表示根据string的前缀来判断以什么进制去解析：0x开头的以16进制的方式去解析，0开头的以8进制方式去解析，其它的以10进制方式解析。
    //func ParseInt(s string, base int, bitSize int) (i int64, err error)
    i, _ := strconv.ParseInt("123", 0,64)
    fmt.Println(i)
    fmt.Println(typeof(i))

    //以5进制方式解析"23"，保存为int64类型：
    i2, _ := strconv.ParseInt("23", 5, 64)
    fmt.Println(i2)
    fmt.Println(typeof(i2))

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
/*
1.234
1.234
123
int64
13
int64
456
789
135
*/
