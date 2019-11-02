/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

示例 2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

示例 4:
输入: [1,3,5,6], 0
输出: 0
*/

package main

import "fmt"

//最好是二分查找,后面实现
/*
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
        var index int = -1
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 && nums[i] < target {
			index = i+1
		}
		if nums[i] == target || nums[i] > target {
			index = i
                        break
		}

	}
        return index
}
*/

//二分查找
func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			if m+1 >= r || nums[m+1] > target {
				return m + 1
			} else {
				l = m + 1
			}
		} else { //nums[m] > target
			if m == 0 || nums[m-1] < target {
				return m
			} else {
				r = m - 1
			}
		}
	}
	return 0
}

func main() {
	var nums []int = []int{1, 3, 5, 6, 7, 9, 10}
	var t1 int = 5
	var t2 int = 2
	var t3 int = 7
	var t4 int = 0
	var t5 int = 12
	fmt.Println(searchInsert(nums, t1))
	fmt.Println(searchInsert(nums, t2))
	fmt.Println(searchInsert(nums, t3))
	fmt.Println(searchInsert(nums, t4))
	fmt.Println(searchInsert(nums, t5))
}
