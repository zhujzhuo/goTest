package main

import (
	"fmt"
	"math"
)

func binarySearch(arr []int, k int) int {
	left, right, mid := 1, len(arr), 0
	for {
		// mid向下取整
		mid = int(math.Floor(float64((left + right) / 2)))
		if arr[mid] > k {
			// 如果当前元素大于k，那么把right指针移到mid - 1的位置
			right = mid - 1
		} else if arr[mid] < k {
			// 如果当前元素小于k，那么把left指针移到mid + 1的位置
			left = mid + 1
		} else {
			// 否则就是相等了，退出循环
			break
		}
		// 判断如果left大于right，那么这个元素是不存在的。返回-1并且退出循环
		if left > right {
			mid = -1
			break
		}
	}
	// 输入元素的下标
	return mid
}

func binarySearch2(sortedArray []int, lookingFor int) int {
	var low int = 0
	var high int = len(sortedArray) - 1
	for low <= high {
		var mid int = low + (high-low)/2
		var midValue int = sortedArray[mid]
		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

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

func binarySearch4(arr []int, k int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		/**
		利用位与（&）提取出两个数相同的部分，利用异或（^）拿出两个数不同的部分的和，相同的部分加上不同部分的和除2即得到两个数的平均值
		异或： 相同得零，不同得1 == 男男等零，女女得零，男女得子
		avg = (a&b)  + (a^b)>>1;
		或者
		avg = (a&b)  + (a^b)/2;
		*/
		mid := low&high + (low^high)>>1
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

func main() {

	sortedArray := []int{1, 3, 5, 6, 7, 8, 10, 30, 40, 44, 45, 56, 79}

	fmt.Println(binarySearch3(sortedArray, 3))

}
