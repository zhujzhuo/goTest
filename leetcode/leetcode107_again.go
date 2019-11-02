/*
107:
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/

package main

import (
	"fmt"
)

//递归实现
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := new(Queue)
	queue.Put(root)
	for queue.Len() > 0 {
		vals := make([]int, 0)
		result = append(result, )
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Poll()
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue.Put(node.Left)
			}
			if node.Right != nil {
				queue.Put(node.Right)
			}
		}
		result = append(result, vals)
	}

	for i := 0; i < len(result)/2; i++ {
		j := len(result) - i - 1
		result[i], result[j] = result[j], result[i]
	}

	return result
}

type Queue struct {
	data []*TreeNode
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Put(n *TreeNode) {
	q.data = append(q.data, n)
}

func (q *Queue) Poll() *TreeNode {
	n := q.data[0]
	q.data = q.data[1:]
	return n
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

	fmt.Println(levelOrderBottom(&t1))       // 2 2 1
	fmt.Println(levelOrderBottom(&t2))       // 4 3 3 4 2 2 1
	fmt.Println(levelOrderBottom(&t3))       // 2 1

}
