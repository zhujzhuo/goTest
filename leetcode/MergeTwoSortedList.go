package main

import (
"fmt"
)
// ListNode define
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

func  main(){
      	//创建有序链表2——>3——>4
		var v3 = &ListNode{10, nil}
		var v2 = &ListNode{5, v3}
		var l1 = ListNode{2, v2}
    
		//创建有序链表4——>5——>6
		var m3 = &ListNode{20, nil}
		var m2 = &ListNode{7, m3}
		var l2 = ListNode{3, m2}
 

		var l3 *ListNode = mergeTwoLists(&l1, &l2)
		for l3.Next != nil {
			fmt.Printf("%d——>", l3.Val)
			l3 = l3.Next
		}
		//打印最后一个值
		fmt.Printf("%d\n", l3.Val)
}