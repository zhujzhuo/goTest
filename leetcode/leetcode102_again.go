/*
102:
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
*/

package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
//102  可以和107 对比学习
//递归实现
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		count := len(queue)
		var seq []int
		for i := 0; i < count; i++ {
			node := queue[i]
			seq = append(seq, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[count:]
		res = append(res, seq)
	}
	return res
*/
//广度优先搜索
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	type BFS struct {
		Node  *TreeNode
		Depth int
	}
	q := []BFS{
		{root, 0},
	}
	var ret [][]int
	for len(q) > 0 {
		c := q[0]
		node := c.Node
		q = q[1:]
		if c.Depth == len(ret) {
			ret = append(ret, []int{node.Val})
		} else {
			p := ret[c.Depth]
			p = append(p, node.Val)
			ret[c.Depth] = p
		}
		if node.Left != nil {
			q = append(q, BFS{
				Node:  node.Left,
				Depth: c.Depth + 1,
			})
		}
		if node.Right != nil {
			q = append(q, BFS{
				Node:  node.Right,
				Depth: c.Depth + 1,
			})
		}
	}
	return ret
}

func main() {
	var n3 = &TreeNode{2, nil, nil}
	var n2 = &TreeNode{2, nil, nil}
	var t1 = TreeNode{1, n2, n3}

	var m7 = &TreeNode{4, nil, nil}
	var m6 = &TreeNode{3, nil, nil}
	var m5 = &TreeNode{3, nil, nil}
	var m4 = &TreeNode{4, nil, nil}
	var m3 = &TreeNode{2, m6, m7}
	var m2 = &TreeNode{2, m4, m5}
	var t2 = TreeNode{1, m2, m3}

	var o2 = &TreeNode{2, nil, nil}
	var t3 = TreeNode{1, o2, nil}

	fmt.Println(levelOrder(&t1)) // 1 2 2
	fmt.Println(levelOrder(&t2)) // 1 2 2 4 3 3 4
	fmt.Println(levelOrder(&t3)) // 1 2

}
