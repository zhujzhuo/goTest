package main

/*
new 分配
Go提供了两种分配原语，即内建函数 new 和 make。 它们所做的事情不同，所应用的类型也不同。它们可能会引起混淆，但规则却很简单。 让我们先来看看 new。这是个用来分配内存的内建函数， 但与其它语言中的同名函数不同，它不会初始化内存，只会将内存置零。 也就是说，new(T) 会为类型为 T 的新项分配已置零的内存空间， 并返回它的地址，也就是一个类型为 *T 的值。用Go的术语来说，它返回一个指针， 该指针指向新分配的，类型为 T 的零值。

既然 new 返回的内存已置零，那么当你设计数据结构时， 每种类型的零值就不必进一步初始化了，这意味着该数据结构的使用者只需用 new 创建一个新的对象就能正常工作。例如，bytes.Buffer 的文档中提到“零值的 Buffer 就是已准备就绪的缓冲区。" 同样，sync.Mutex 并没有显式的构造函数或 Init 方法， 而是零值的 sync.Mutex 就已经被定义为已解锁的互斥锁了。

“零值属性”可以带来各种好处。考虑以下类型声明。

type SyncedBuffer struct {
	lock    sync.Mutex
	buffer  bytes.Buffer
}
SyncedBuffer 类型的值也是在声明时就分配好内存就绪了。后续代码中， p 和 v 无需进一步处理即可正确工作。

p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer

再回到内存分配上来。内建函数 make(T, args) 的目的不同于 new(T)。它只用于创建切片、映射和信道，并返回类型为 T（而非 *T）的一个已初始化 （而非置零）的值。 出现这种用差异的原因在于，这三种类型本质上为引用数据类型，它们在使用前必须初始化。 例如，切片是一个具有三项内容的描述符，包含一个指向（数组内部）数据的指针、长度以及容量， 在这三项被初始化之前，该切片为 nil。对于切片、映射和信道，make 用于初始化其内部的数据结构并准备好将要使用的值。例如，

make([]int, 10, 100)
会分配一个具有100个 int 的数组空间，接着创建一个长度为10， 容量为100并指向该数组中前10个元素的切片结构。（生成切片时，其容量可以省略，更多信息见切片一节。） 与此相反，new([]int) 会返回一个指向新分配的，已置零的切片结构， 即一个指向 nil 切片值的指针。

下面的例子阐明了 new 和 make 之间的区别：

var p *[]int = new([]int)       // 分配切片结构；*p == nil；基本没用
var v  []int = make([]int, 100) // 切片 v 现在引用了一个具有 100 个 int 元素的新数组

// 没必要的复杂：
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// 习惯用法：
v := make([]int, 100)
var v  []int = make([]int, 100)

请记住，make 只适用于映射、切片和信道且不返回指针。若要获得明确的指针， 请使用 new 分配内存。
*/
import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1   = Vertex{1, 2} // 类型为 Vertex
	v2   = Vertex{X: 1} // Y:0 被省略
	v3   = Vertex{}     // X:0 和 Y:0
	pvar = &v1          // 类型为 *Vertex
)

func main() {
	fmt.Println(v1, pvar.X, v2, v3)
	v := Vertex{1, 2}
	fmt.Println(v)
	pv := &v
	v.X = 4
	pv.Y = 1e3
	fmt.Println(v)
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j // point to j
	fmt.Println(*p)
	fmt.Println(&p)
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
