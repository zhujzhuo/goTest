/*
flag  重点（命令行格式的需要用到很多）
文件操作：fileget   find_file   操作系统文件的读写，目录的遍历等
自定义error : errors    stringer.go（自定义格式非常重要）
指针:pointer
sync_exercise  可以和channel 一起找并发的实例。可以看具体的代码，首先看看gh-ost的代码，具体实现一下功能
zap 重点（日志是每个系统必须的）

godoc 既是一个程序，又是一个Web服务器，它对Go的源码进行处理，并提取包中的文档内容。 出现在顶级声明之前，
且与该声明之间没有空行的注释，将与该声明一起被提取出来，作为该条目的说明文档。 这些注释的类型和风格决定了 
godoc 生成的文档质量
每个包都应包含一段包注释，即放置在包子句前的一个块注释。对于包含多个文件的包， 包注释只需出现在其中的任一文件中即可。
包注释应在整体上对该包进行介绍，并提供包的相关信息。 它将出现在 godoc 页面中的最上面，并为紧随其后的内容建立详细的文档。
在包中，任何顶级声明前面的注释都将作为该声明的文档注释。 在程序中，每个可导出（首字母大写）的名称都应该有文档注释。

// Compile 用于解析正则表达式并返回，如果成功，则 Regexp 对象就可用于匹配所针对的文本。
func Compile(str string) (regexp *Regexp, err error) {

*/
package main
//import "math"
import "fmt"
/*
1. 点操作
我们有时候会看到如下的方式导入包
import(
. "fmt"
)
这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调
用的fmt.Println("hello world")可以省略的写成Println("hello world")
2. 别名操作
别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
import(
f "fmt"
)
别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")
3. _操作
这个操作经常是让很多人费解的一个操作符，请看下面这个import
import (
"database/sql"
_ "github.com/ziutek/mymysql/godrv"
)
_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。

*/


var v1 int=10
// v1:=10
// var v1=10
var v2 string 
var v3 [10]int  // 数组
var v4 []int   // 数组切片
var v5 struct {
f int }
var v6 *int   // 指针
var v7 map[string]int   // map,key为string类型,value为int类型

var v8 func(a int) int

var (
v9 int
v10 string
)


const Pi float64 = 3.14159265358979323846 
const zero = 0.0   // 无类型浮点常量 
const (
size int64 = 1024 
eof = -1   // 无类型整型常量
)

const u, v float32 = 0, 3   // u = 0.0, v = 3.0,常量的多重赋值
const a, b, c = 3, 4, "foo" //a=3,b=4,c="foo", 无类型整型和字符串常量

//在包中，任何顶级声明前面的注释都将作为该声明的文档注释。 在程序中，每个可导出（首字母大写）的名称都应该有文档注释。 
//浮点数的比较
// p为用户自定义的比较精度,比如0.00001 
//func IsEqual(f1, f2, p float64) bool {
//   return math.Fdim(f1, f2) < p 
//}

[32]byte // 长度为32的数组,每个元素为一个字节 
[2*N] struct { x, y int32 } // 复杂类型数组
array := [5]int{1,2,3,4,5}
[1000]*float64  // 指针数组
[3][5]int  // 二维数组
[2][2][2]float64   // 等同于[2]([2]([2]float64))
var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
  fmt.Println("Hello World!")	
  fmt.Println(a)
  var  str1  string ="hello"
  var  str2  string ="123"
  fmt.Println(str1+str2)
  fmt.Println(len(str1))
}

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

//注意点：
//常量定义的右值也可以是一个在编译期运算的常量表达式,比如
//const mask = 1 << 3 由于常量的赋值是一个编译期行为,所以右值不能出现任何需要运行期才能得出结果的表达式

// iota在每个const开头被重设为0
// iota比较特殊,可以被认为是一个可被编译器修改的常量,在每一个const关键字出现时被重置为0,
//然后在下一个const出现之前,每出现一次iota,其所代表的数字会自动增1
//如果两个const的赋值语句的表达式是一样的，那么可以简写后一个赋值表达式。
const(
   c0 = iota   //0
   c1          //1
   c2          //2
)
const(
   a0 = 1 <<iota  //1
   a1             //2
   a2             //4
)



// 以大写字母开头的常量在包外可见。const(Num1 Num2 numberOfDays),以上例子中numberOfDays为包内私有,其他符号则可被其他包访问

//两个不同类型的整型数不能直接比较,比如int8类型的数和int类型的数不能直接比较
//var i int64
//var j int32
//i,j=1,2
//if i==j { } //编译会报错



//Go没有逗号操作符，而 ++ 和 -- 为语句而非表达式。 因此，若你想要在 for 
//中使用多个变量，应采用平行赋值的方式 （因为它会拒绝 ++ 和 --）
//反转 a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	a[i], a[j] = a[j], a[i]
}


//switch 并不会自动下溯，但 case 可通过逗号分隔来列举相同的处理条件。
func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}


//下面的例子阐明了 new 和 make 之间的区别：
var p *[]int = new([]int)       // 分配切片结构；*p == nil；基本没用
var v  []int = make([]int, 100) // 切片 v 现在引用了一个具有 100 个 int 元素的新数组

// 没必要的复杂：
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// 习惯用法：
v := make([]int, 100)
//请记住，make 只适用于映射、切片和信道且不返回指针。若要获得明确的指针， 请使用 new 分配内存。
//new(T) 会为类型为 T 的新项分配已置零的内存空间， 并返回它的地址，也就是一个类型为 *T 的值。用Go的术语来说，它返回一个指针， 该指针指向新分配的，类型为 T 的零值。

