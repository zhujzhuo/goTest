package main
//flag 使用具体说明  https://github.com/astaxie/gopkg/tree/master/flag
//go run  flag.go  -deltaT 61m,72h,80s
import (
  "errors"
  "flag"
  "fmt"
  "strings"
  "time"
)

type interval []time.Duration

/*
type ee int
还可以创建自定义 flag类型（int bool等都是类型），只要实现 flag.Value 接口即可
// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
type Value interface {
    String() string
    Set(string) error
}
所有参数类型需要实现 Value 接口，flag 包中，为int、float、bool等实现了该接口。借助该接口，我们可以自定义flag。

*/

//实现String接口
func (i *interval) String() string {
  return fmt.Sprintf("intervalFlag: %v", *i)
}

//实现Set接口,Set接口决定了如何解析flag的值
func (i *interval) Set(value string) error {
  //此处决定命令行是否可以设置多次-deltaT
  if len(*i) > 0 {
    return errors.New("interval flag already set")
  }
  for _, dt := range strings.Split(value, ",") {
    duration, err := time.ParseDuration(dt)
    if err != nil {
      return err
    }
    *i = append(*i, duration)
  }
  return nil
}

var intervalFlag interval

func init() {
  flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
  flag.Parse()
  fmt.Println(intervalFlag)
}
