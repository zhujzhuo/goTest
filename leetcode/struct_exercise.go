package main

import "fmt"

type Person struct {
   name string
   sex  string
   age int
}
func main() {
   person1 := Person{"zhangsan","man",25} //创建一个person1对象
   fmt.Printf("person1:%v\n",person1)
   demo(&person1)
   fmt.Printf("person1:%v\n",person1)
}

func demo(person *Person)  {
   (*person).age = 18 //显示的解引用
   person.name = "GoLang" //隐式的解引用
}
