package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
/* 
声明变量，默认 map 是 nil 
var map_variable map[key_data_type]value_data_type

使用 make 函数 
map_variable := make(map[key_data_type]value_data_type)
*/
//map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值
//map 的文法跟结构体文法相似，不过必须有键名
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
// 如果顶级的类型只有类型名的话，可以在文法的元素中省略键名
var m1 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

var v []int 
func main() {
	fmt.Println(m)
	fmt.Println(m1)
	m2 := make(map[string]int)

	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])

	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])
        //如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
        //同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
