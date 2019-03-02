package main
import (
    "fmt"
    "flag"
)

func main(){
//flag提供了Arg(i),Args()来获取non-flag参数，NArg()来获取non-flag的个数
//flag还提供了NFlag()获取那些匹配上的参数的个数
    data_path := flag.String("D","/home/manu/sample/","DB data path")
    log_file := flag.String("l","/home/manu/sample.log","log file")
    nowait_flag :=flag.Bool("W",false,"do not wait until operation completes")

    flag.Parse()

    var cmd string = flag.Arg(0);

    fmt.Printf("action   : %s\n",cmd)
    fmt.Printf("data path: %s\n",*data_path)
    fmt.Printf("log file : %s\n",*log_file)
    fmt.Printf("nowait   : %v\n",*nowait_flag)

    fmt.Printf("-------------------------------------------------------\n")

    fmt.Printf("there are %d non-flag input param\n",flag.NArg())
    fmt.Printf("there are %d flag input param\n",flag.NFlag())
    for i,param := range flag.Args(){
        fmt.Printf("#%d    :%s\n",i,param)
    }


}
