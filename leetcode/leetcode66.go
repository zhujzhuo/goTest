/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:
输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。

示例 2:
输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

示例 3:
输入: [9,9,9,9]
输出: [1,0,0,0,0]
解释: 输入数组表示数字 9999。
*/

package main

import "fmt"

func plusOne(digits []int) []int {
	len_di := len(digits)
	var res []int
	// 只有全部元素是9的时候结果数组会多一位，否则位数不变
	for i := len_di - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			break
		}
	}
	//表示全部是9的情况，第一位才会是0,结果集的其它位全部是0
	if digits[0] == 0 {
		//利用make初始化其它位数为0，长度为len+1
		res = make([]int, len_di+1)
		res[0] = 1
	} else {
		res = digits
	}

	return res
}
func main() {
	var nums1 []int = []int{9, 9, 9, 9}
	var nums2 []int = []int{9, 9}
	var nums3 []int = []int{1, 2, 3, 4}
	var nums4 []int = []int{1, 2, 3, 9}
	fmt.Println(plusOne(nums1))
	fmt.Println(plusOne(nums2))
	fmt.Println(plusOne(nums3))
	fmt.Println(plusOne(nums4))

}

/*
func plusOne(digits []int) []int {
  carry := 1
  res := make([]int, len(digits)+1)
  for i := len(digits) -1 ; i >= 0; i-- {
    if i == len(digits) -1 {
      tmp := digits[i]
      digits[i] = (digits[i] + carry ) % 10
      carry =  (tmp + carry) / 10
      res[i+1] = digits[i]
      res[i] = carry
    }else {
      tmp := digits[i]
      digits[i] = (tmp + carry) % 10
      carry = (tmp +carry )/ 10
      res[i+1] = digits[i]
      res[i] = carry
    }
  }
  i := 0
  for i < len(digits) && res[i] == 0 {
    i++
  }
  return res[i:]
}
*/
