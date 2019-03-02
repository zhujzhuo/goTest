/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
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
 
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var result int = 0

//二叉树深度递归找左右叶子节点的最大深度
//后面实现按照层级遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		lefthDepth := maxDepth(root.Left) + 1
		rightDepth := maxDepth(root.Right) + 1
		result = Max(lefthDepth, rightDepth)
	}
	return result
}

func main() {
       var n3 = &TreeNode{2,nil,nil}
       var n2 = &TreeNode{2,nil,nil}
       var t1 = TreeNode{1,n2,n3}

       var m7 = &TreeNode{4,nil,nil}
       var m6 = &TreeNode{3,nil,nil}
       var m5 = &TreeNode{3,nil,nil}
       var m4 = &TreeNode{4,nil,nil}
       var m3 = &TreeNode{2,m6,m7}
       var m2 = &TreeNode{2,m4,m5}
       var t2 = TreeNode{1,m2,m3}

       var o2 = &TreeNode{2,nil,nil}
       var t3 = TreeNode{1,o2,nil}

       fmt.Println(Min(1,2)) //1     
       fmt.Println(maxDepth(&t1))  //2
       fmt.Println(maxDepth(&t2))  //3
       fmt.Println(maxDepth(&t3))  //2

}
