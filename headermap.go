package main
// Header类型用于表达HTTP头部的键值对信息 
type Header map[string][]string
// Add()方法用于添加一个键值对到HTTP头部
// 如果该键已存在,则会将值添加到已存在的值后面 
func (h Header) Add(key, value string) {
    textproto.MIMEHeader(h).Add(key, value)
}
// Set()方法用于设置某个键对应的值,如果该键已存在,则替换已存在的值 
func (h Header) Set(key, value string) {
    textproto.MIMEHeader(h).Set(key, value)
}
// 还有更多其他方法
// done played first
type Integer int
func (a Integer) Less(b Integer) bool { 
	return a < b
 }

func main() {
var a Integer = 1 if a.Less(2) {
            fmt.Println(a, "Less 2")
        }
 }

func (a *Integer) Add(b Integer) { 
	*a += b
 }
func (a Integer) Add(b Integer) { 
 	a += b
}
