package main 


import (
    "strconv"
    "fmt"
    "time"
)

func EscapeName(name string) string {
	if unquoted, err := strconv.Unquote(name); err == nil {
		name = unquoted
	}
        fmt.Println(name)
	return fmt.Sprintf("`%s`", name)
}

func main(){
    str := "table_name_test"
    unquoted, _ := strconv.Unquote(`"1table_name_test"`)
    fmt.Println(unquoted)
    fmt.Println(EscapeName(str))
    list,_ := strconv.Unquote(`"Hello\t\u4e16\u754c\uff01"`)
    fmt.Println(list)
    fmt.Println(time.Now())
    fmt.Println(rune(0))
    
}
