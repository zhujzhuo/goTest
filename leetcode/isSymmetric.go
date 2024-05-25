/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
//将root赋值给两个树，递归比对子节点的左侧（右侧）和另外一个子节点的右侧（左侧）相同与否
func isSymmetric(root *TreeNode) bool {
     return isMirror(root,root)
}
func isMirror(p *TreeNode,q *TreeNode) bool {
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
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	return isSame(root.Left, root.Right)

}

func isSame(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if q != nil && p != nil && p.Val == q.Val {
		return isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
	} else {
		return false
	}
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

	fmt.Println(isSymmetric(&t1)) //true
	fmt.Println(isSymmetric(&t2)) //true
	fmt.Println(isSymmetric(&t3)) //false
}
