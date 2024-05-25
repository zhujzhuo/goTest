/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//使用递归的方法
//列表合并，每次去找到最小，赋值给res，Next的值就是下次递归取得的最小值
/*
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	//如果有一条链是nil，直接返回另外一条链
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	//当l1节点的值大于l2节点的值，那么res指向l2的节点，从l2开始遍历，反之从l1开始
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
}
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		return &ListNode{l1.Val, mergeTwoLists(l1.Next, l2)}
	} else {
		return &ListNode{l2.Val, mergeTwoLists(l1, l2.Next)}
	}
}
func main() {
	//创建有序链表1——>4——>5
	var v3 = &ListNode{5, nil}
	var v2 = &ListNode{4, v3}
	var l1 = ListNode{1, v2}

	//创建有序链表2——>6
	var m2 = &ListNode{6, nil}
	var l2 = ListNode{2, m2}

	//创建有序链表1——>6——>9——>10
	var t4 = &ListNode{10, nil}
	var t3 = &ListNode{9, t4}
	var t2 = &ListNode{6, t3}
	var l3 = ListNode{1, t2}

	//restmp 是一个*ListNode
	var restmp = mergeTwoLists(&l1, &l2)
	var res = mergeTwoLists(restmp, &l3)
	for res.Next != nil {
		fmt.Printf("%d——>", res.Val)
		res = res.Next
	}
	//打印最后一个值
	fmt.Printf("%d\n", res.Val)
}
