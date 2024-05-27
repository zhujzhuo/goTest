package main
import "fmt"
//import "errors"
import "mymath"

type PersonInfo  struct {
	Id string
	Name string
	Address string
}

func main(){
    var personDB map[string] PersonInfo
    //personDB是声明的map变量名,string是键的类型,PersonInfo则是其中所存放的值类型
    personDB = make(map[string] PersonInfo,100)
    //初始化
    // 往这个map里插入几条数据
    personDB["1234"] = PersonInfo{"12345", "Tom", "Room 203,..."} 
    personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
    // 从这个map查找键为"1234"的信息 
    person, ok := personDB["1234"]
    // ok是一个返回的bool型,返回true表示找到了对应的数据 
    if ok {
            fmt.Println("Found person", person.Name, "with ID 1234.")
    } else {
            fmt.Println("Did not find person with ID 1234.")
    }
    c ,_:= mymath.Add(1,2)  //mymath包内函数调用
    fmt.Println(c)
}

