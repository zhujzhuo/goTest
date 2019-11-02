/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 正常的按照位数相加，保存加和的个位数remainder，保存进位carry ,进位加到下次的两数加和中,
// 两数加和最大是9+9+1=19（最大数+进位），因此，carry只能为0或者1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	var x, y int //为了解决位数不同的数字加和，将多出来的位数置为0和另外一个做加和(较小数的高位补齐为0)
	var carry int = 0
	var remainder int = 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next // 此处赋值避免在l1是nil的时候还去给它赋值Next，会有null pointer 报错
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}
		remainder = (x + y + carry) % 10
		carry = (x + y + carry) / 10 //进位只能是0或者1，也可以用x + y + carry 是否大于9来判断0 1
		node := &ListNode{remainder, nil}
		if head == nil { //第一次赋值,以后的赋值只需要移动tail即可
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}

	}
	//最后一位进位处理
	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}
	return head
}

func main() {
	/*
		//创建有序链表2——>4——>3
		var v3 = &ListNode{3, nil}
		var v2 = &ListNode{4, v3}
		var l1 = ListNode{2, v2}

		//创建有序链表5——>6——>4
		var m3 = &ListNode{4, nil}
		var m2 = &ListNode{6, m3}
		var l2 = ListNode{5, m2}
	*/

	//创建有序链表9——>9
	var v2 = &ListNode{9, nil}
	var l1 = ListNode{9, v2}

	//创建有序链表9——>9
	var m2 = &ListNode{9, nil}
	var l2 = ListNode{9, m2}

	var l3 *ListNode = addTwoNumbers(&l1, &l2)
	for l3.Next != nil {
		fmt.Printf("%d——>", l3.Val)
		l3 = l3.Next
	}
	//打印最后一个值
	fmt.Printf("%d\n", l3.Val)

}

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if (l2 == nil) {
        return l1
    }
    l := new(ListNode)
    next := l
    flag := false //是否有进位的表示
    for l1 != nil && l2 != nil {
        val := l1.Val + l2.Val
        if flag {
            val += 1
        }

        if val >= 10 {
            flag = true
        } else {
            flag = false
        }

        node := &ListNode {
            val % 10,
            nil,
        }

        next.Next = node
        next = node
        l1 = l1.Next
        l2 = l2.Next
    }

    var l3 *ListNode
    if l1 == nil {
        l3 = l2
    } else {
        l3 = l1
    }

    for l3 != nil {
        val := l3.Val
        if flag {
            val += 1
        } else {
            next.Next = l3
            break
        }

        if val >= 10 {
            flag = true
        } else {
            flag = false
        }

        node := &ListNode {
            val % 10,
            nil,
        }
        next.Next = node
        next = node
        l3 = l3.Next
    }

    if flag {
        node := &ListNode {
            1,
            nil,
        }
        next.Next = node
    }

    return l.Next
}
*/
