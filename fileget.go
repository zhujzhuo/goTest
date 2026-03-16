package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dir_list, e := ioutil.ReadDir("/Users/didi/gocode/src/github.com/dddddd")
	if e != nil {
		fmt.Printf("read dir error:%s\n", e)
		return
	}
	var filename string
	for i, v := range dir_list {
		filename = v.Name()
		fmt.Println(i, "=", filename)
	}
}
