package main

import "fmt"

func binarySearch(arr []int, k int) int {
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

func mergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		return merge(mergeSort(list[:len(list)/2]), mergeSort(list[len(list)/2:]))
	}
}

//合并两个有序列表到一个列表中
func merge(list0 []int, list1 []int) []int {
	var result []int
	index0 := 0
	index1 := 0
	for {
		if index0 < len(list0) && index1 < len(list1) {
			if list0[index0] < list1[index1] {
				result = append(result, list0[index0])
				index0 = index0 + 1
			} else {
				result = append(result, list1[index1])
				index1 = index1 + 1
			}
		} else {
			break
		}
	}
	if index0 < len(list0) {
		result = append(result, list0[index0:]...)
	}
	if index1 < len(list1) {
		result = append(result, list1[index1:]...)
	}
	return result
}

func twoSum(nums []int, target int) []int {
	tables := make(map[int]int)
	for index, value := range nums {
		tmp := target - value
		if i, ok := tables[tmp]; ok {
			return []int{i, index}
		} else {
			//tables 键值对是数组中的 值（value）和下标 （index）
			//比如3  4 target=7 ，搜到3的时候tables里面还没有4，因此将3加入进去，当搜到4的时候，3已经在tables中了，返回即可
			//这也就是为什么这里不先遍历一次数组全部赋值给map的原因，最坏的情况下就是最后两个元素符合条件，最好情况是第一二个元素符合条件
			//最多遍历一遍，最少只需要查询前两个元素
			tables[value] = index
		}
	}
	return []int{}
}


//二叉树深度递归找左右叶子节点的最大深度
//后面实现按照层级遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		//lefthDepth := maxDepth(root.Left) + 1
		//rightDepth := maxDepth(root.Right) + 1
		//result = Max(lefthDepth, rightDepth)
		return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
	}
}

//左右两边对称
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		} else {
			return false
		}

	} else {
		return false
	}
}

//左右两边镜像
func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {

	sortedArray := []int{1, 3, 5, 6, 7, 8, 10, 30, 40, 44, 45, 56, 79}
	array := []int{4, 7, 5, 10, 3, 8, 20, 13, 45, 65, 12, 15, 25}

	for i, value := range array[:3] {
		fmt.Printf("index:%v  value:%v\n", i, value)
	}
	fmt.Println(binarySearch(sortedArray, 6))

	var sortarray []int = mergeSort(array)
	for i, value := range sortarray {
		fmt.Printf("index:%v  value:%v\n", i, value)
	}

	bytemap := map[byte]byte{'{': '}', '(': ')', '"': '"'}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k, v := range bytemap {
		fmt.Printf("%v -> %v\n", k, v)
	}

	var n3 = &TreeNode{2, nil, nil}
	var n2 = &TreeNode{2, nil, nil}
	var t1 = TreeNode{1, n2, n3}

	fmt.Println(maxDepth(&t1)) //2
}
