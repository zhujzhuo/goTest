/*
go run  flag_definevalue.go  -slice="java,go,php     scr"  -testint=3307  -flagname=123
flag.Var
flag.Intvar
flag.Int

go run  flag_definevalue.go  -slice="java,go,php     scr"  -testint=3307  -flagname=123 -name="test a string var"

*/
package main

import (
	"flag"
	"fmt"
	"strings"
)

//定义一个类型，用于增加该类型方法
type sliceValue []string

//new一个存放命令行参数值的slice
func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

/*
Value接口：
type Value interface {
    String() string
    Set(string) error
}
实现flag包中的Value接口，将命令行接收到的值用“,”分隔存到slice里
*/
func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

//flag为slice的默认值default is me,和return返回值没有关系
func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("default is me", ","))
	return "It's none of my business"
}

/*
可执行文件名 -slice="java,go"  最后将输出[java,go]
可执行文件名 最后将输出[default is me]
*/
/*
可以通过如下方式定义该 flagl类型：

flag.Var(&flagVal, "name", "help message for flagname")
*/

func main() {
	var languages []string
	var flagVal []string
	var Intv int
        var c string
	flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
	flag.Var(newSliceValue([]string{}, &flagVal), "name", "help message for flagname")
        
        //另一种方式
        q := flag.Bool("q", false, "suppress non-error messages during configuration testing")
	ip := flag.Int("flagname", 1234, "help message for flagname") //with type *int
	flag.IntVar(&Intv, "testint", 3306, "test int var `port`")           //&Intv
        flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.Parse()

	//打印结果slice接收到的值
	fmt.Println(languages)
	fmt.Println(Intv)
	fmt.Println(flagVal)
	fmt.Println(*ip)
	fmt.Println(*q)
	fmt.Println(c)
        flag.PrintDefaults()
}
