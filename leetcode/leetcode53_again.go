/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

进阶:
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
package main

import "fmt"

/*
定义一个函数f(n)，以第n个数为结束点的子数列的最大和，存在一个递推关系f(n) = max(f(n-1) + nums[n],f(n-1));
将这些最大和保存下来后，取最大的那个就是最大子数组和。因为最大连续子数组等价于 最大的以n个数为结束点的子数列和
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
        // 一段一段的子数列的和
	res := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
                // 前面累加的增长 > 0 那就加上我自己再看看
		if sum > 0 {
			sum = sum + nums[i]
		} else {
                // 前面累加的增长 < 0  前面的所有累加的增长都白干了，那就从这一次开始看后面的吧
			sum = nums[i]
		}
               // sum就等于从前面某一天到今天的增长
		if sum > res {
			res = sum
		}
	}
	return res
}

func main() {
	var nums []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	var nums2 []int = []int{-2, -3, -1, -5}
	fmt.Println(maxSubArray(nums))
	fmt.Println(maxSubArray(nums2))
}
