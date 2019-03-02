/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
*/

package main

import "fmt"

/*
1、将nums2赋值到nums1的index=m之后的位置，调用排序即可
2、倒序扫描两个数组，并且nums[i]和nums[j] 相等时,移动i,因为后面的元素不涉及元素移动
   正向扫描要移动已有的元素
*/
func merge(nums1 []int, m int, nums2 []int, n int) []int{
	if len(nums1) < m+n {
		return []int{}
	}
	i, j, index := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[index] = nums2[j]
			j--
			index--
		} else {
			nums1[index] = nums1[i]
			i--
			index--
		}
	}
	if j >= 0 {
		for j >= 0 {
			nums1[j] = nums2[j]
                        j--
		}
	}
	return nums1
}

func main() {
      var nums1 []int = []int{2,4,5,7,0,0,0,0,0}
      var nums2 []int = []int{1,3,5,6,7}
      fmt.Println(merge(nums1,4,nums2,5))    
      
}
