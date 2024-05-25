/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:
给定数组 nums = [1,1,2],
函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,0,1,1,1,2,2,3,3,4],
函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
你不需要考虑数组中超出新长度后面的元素。
说明:
为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
你可以想象内部操作如下:
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);
// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/

package main

import "fmt"

/*
数组完成排序后，我们可以放置两个指针i和 j，其中i 是慢指针,而j是快指针。只要nums[i] = nums[j]，我们就增加j以跳过重复项。
当我们遇到 nums[j] != nums[i]时，跳过重复项的运行已经结束，因此我们必须把nums[j]的值复制到nums[i+1]。然后递增i，
接着我们将再次重复相同的过程，直到j 到达数组的末尾为止。
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

/*
func removeDuplicates(nums []int) int {
	idx := 0
	for i, v := range nums {
		if i == 0 || v == nums[idx] {
			continue
		}
		if v < nums[idx] {
			break
		}
		nums[idx+1] = v
		idx++
	}
	return idx + 1
}
*/
func main() {

	var nums1 []int = []int{1, 2, 3, 4, 4, 4, 5, 7, 7, 8, 9}
	var nums2 []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	var nums3 []int = []int{0, 1, 1, 2, 2, 3, 4, 4, 5}
	fmt.Println(removeDuplicates(nums1))
	fmt.Println(nums1)

	fmt.Println(removeDuplicates(nums2))
	fmt.Println(nums2)

	fmt.Println(removeDuplicates(nums3))
	fmt.Println(nums3)
}
