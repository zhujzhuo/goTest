/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
package main

import (
    "fmt"
    "time"
)

//可以赋值到map中遍历一遍即可,但是内存占用就会多
func twoSum(nums []int, target int) []int {
    tables := make(map[int]int)
    for index,value := range nums {
        tmp := target - value
        if i,ok := tables[tmp]; ok {
            return []int{i,index}
        }else{
        //tables 键值对是数组中的 值（value）和下标 （index）
        //比如3  4 target=7 ，搜到3的时候tables里面还没有4，因此将3加入进去，当搜到4的时候，3已经在tables中了，返回即可
        //这也就是为什么这里不先遍历一次数组全部赋值给map的原因，最坏的情况下就是最后两个元素符合条件，最好情况是第一二个元素符合条件
        //最多遍历一遍，最少只需要查询前两个元素
            tables[value] = index
        }    
    }
    return []int{}
}
/*
//尽量减少中间结果的赋值，减少内存，
//每个元素的target和后面所有的元素最比较
func twoSum(nums []int, target int) []int {
	for i := 0 ; i < len(nums); i++ {
            for j:=i+1; j < len(nums); j++ {
               if nums[j] == target - nums[i] {
                  return []int{i,j}
               }
            }
	}
        return []int{}
}
*/
func main(){
     starttime := time.Now()
     nums := []int{2, 7, 11, 15}
     target := 9
     fmt.Println(twoSum(nums,target))
     fmt.Printf("using time:%v",time.Since(starttime))
     fmt.Println()
}
