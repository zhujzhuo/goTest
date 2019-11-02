package main

import (
	"fmt"
)

//二分查找
func binarySearch3(arr []int, k int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		// 这种写法防止两数和导致的内存溢出
		mid := low + (high-low)>>1 // avg=(a+b)>>1://右移表示除2，左移表示乘2
		if k < arr[mid] {
			high = mid - 1
		} else if k > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//插值查找
//折半查找的进化版，自适应中间值 根据 (关键值 - 起始值) / (末位值 - 起始值) 的比例来决定中间值的下标
func insertSearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	time := 0
	for low < high {
		time += 1
		// 计算mid值是插值算法的核心代码 实际上就是
		mid := low + int((high-low)*(key-arr[low])/(arr[high]-arr[low]))
		if key < arr[mid] {
			high = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {

	sortedArray := []int{1, 3, 5, 6, 7, 8, 10, 30, 40, 44, 45, 56, 79}
	fmt.Println(binarySearch3(sortedArray, 3))
	fmt.Println(insertSearch(sortedArray, 3))

}
