// in example.go
package example

var start int

func Add(n int) int {
    start += n
    return start
}
func IntMin(a, b int) int {
        if a < b {
                return a
        } else {
                return b
        }
}
