package main

//传入一个有序的数组，得到绝对值最大或者最小的元素
import (
	"fmt"
	"math"
)

func getminMbs(arr []int) int {
	len_arr := len(arr)
	min_num := arr[0]
	max_num := arr[len_arr-1]
	if min_num >= 0 {
		return min_num
	} else if max_num <= 0 {
		return max_num
	} else {
		for _, value := range arr {
			if value <= 0 {
				min_num = value
			} else {
				max_num = value
			}
		}
		tmpmin := math.Abs(float64(min_num))
		tmpmax := math.Abs(float64(max_num))
		if tmpmin < tmpmax {
			return min_num
		} else {
			return max_num
		}

	}
}

func getmaxMbs(arr []int) int {
	len_arr := len(arr)
	min_num := arr[0]
	max_num := arr[len_arr-1]
	if min_num >= 0 {
		return max_num
	} else if max_num <= 0 {
		return min_num
	} else {
		tmpmin := math.Abs(float64(min_num))
		tmpmax := math.Abs(float64(max_num))
		if tmpmin < tmpmax {
			return max_num
		} else {
			return min_num
		}

	}

}

func main() {
	arr1 := []int{-20, -10, -3, -2, 0, 1, 5, 8, 9, 19}
	arr2 := []int{0, 1, 5, 8, 9, 19, 20, 22, 23, 26}
	arr3 := []int{-30, -28, -25, -22, -21, -20, -10, -3, -2, -1}

	fmt.Println(getminMbs(arr1))
	fmt.Println(getminMbs(arr2))
	fmt.Println(getminMbs(arr3))
	fmt.Println(getmaxMbs(arr1))
	fmt.Println(getmaxMbs(arr2))
	fmt.Println(getmaxMbs(arr3))
}
