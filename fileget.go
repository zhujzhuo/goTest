package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dir_list, e := ioutil.ReadDir("/Users/zhujinzhuo/gocode/src/")
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	var filename string
	for i, v := range dir_list {
		filename = v.Name()
		fmt.Println(i, "=", filename)
	}
}
