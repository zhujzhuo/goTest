/*
给定两个二叉树，编写一个函数来检验它们是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:

输入:       1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

输出: true

示例 2:
输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]

输出: false

示例 3:
输入:       1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]
输出: false

*/
package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//递归执行比对即可
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
     if p == nil && q == nil {
        return true
     }else if  p != nil && q != nil{
        if p.Val == q.Val{
           return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
        }else{
           return false
        }
     
     }else{
       return false
     } 
}

func main(){
//初始化TreeNode
       
       var n3 = &TreeNode{2,nil,nil}
       var n2 = &TreeNode{2,nil,nil}
       var t1 = TreeNode{1,n2,n3}
       
       var m3 = &TreeNode{2,nil,nil}
       var m2 = &TreeNode{2,nil,nil}
       var t2 = TreeNode{1,m2,m3}
       
       var o2 = &TreeNode{2,nil,nil}
       var t3 = TreeNode{1,o2,nil}
       
       var p2 = &TreeNode{2,nil,nil}
       var t4 = TreeNode{1,nil,p2}
       
       fmt.Println(isSameTree(&t1,&t2)) //true
       fmt.Println(isSameTree(&t3,&t4)) //false
       fmt.Println(isSameTree(&t1,&t3)) //false
       fmt.Println(isSameTree(&t2,&t4)) //false
}


