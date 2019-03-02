/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
示例 1:
输入: 1->1->2
输出: 1->2

示例 2:
输入: 1->1->2->3->3
输出: 1->2->3
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
      Val int
      Next *ListNode
}

//链表指针的操作可以和leetcode2.go看看

func deleteDuplicates(head *ListNode) *ListNode {
     //简单算法就是从head开始扫描，下一个和当前元素相同，则指向下一个元素的Next，直到下一个元素的Next为nil为止
    var res  *ListNode = head
    for head != nil && head.Next != nil {
        if head.Val == head.Next.Val {
           head.Next = head.Next.Next
        }else{
           head = head.Next
        }
    }
   return res 
}

func main(){
        var res1 *ListNode
        //创建有序链表1——>2——>2——>3——>3
        var v5 = &ListNode{3, nil}
        var v4 = &ListNode{3, v5}
        var v3 = &ListNode{2, v4}
        var v2 = &ListNode{2, v3}
        var l1 = ListNode{1, v2}
        res1 = deleteDuplicates(&l1)
        for res1.Next != nil {
                fmt.Printf("%d——>", res1.Val)
                res1 = res1.Next
        }
        //打印最后一个值
        fmt.Printf("%d\n", res1.Val)

        var res2 *ListNode
        //创建有序链表1——>2——>2——>3
        var m4 = &ListNode{3, nil}
        var m3 = &ListNode{2, m4}
        var m2 = &ListNode{2, m3}
        var l2 = ListNode{1, m2}
        res2 = deleteDuplicates(&l2)
        for res2.Next != nil {
                fmt.Printf("%d——>", res2.Val)
                res2 = res2.Next
        }
        //打印最后一个值
        fmt.Printf("%d\n", res2.Val)
        
        var res3 *ListNode
        //创建有序链表1——>1——>1——>1
        var t4 = &ListNode{1, nil}
        var t3 = &ListNode{1, t4}
        var t2 = &ListNode{1, t3}
        var l3 = ListNode{1, t2}
        res3 = deleteDuplicates(&l3)
        for res3.Next != nil {
                fmt.Printf("%d——>", res3.Val)
                res3 = res3.Next
        }
        //打印最后一个值
        fmt.Printf("%d\n", res3.Val)
}

