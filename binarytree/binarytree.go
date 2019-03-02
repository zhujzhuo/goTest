// binary_tree 二叉树
package main

import (
	"fmt"
	"reflect"
)

// 二叉树定义
type BinaryTree struct {
	Data   interface{}
	Lchild *BinaryTree
	Rchild *BinaryTree
}

// 构造方法
func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{Data: data}
}

// 先序遍历
func (bt *BinaryTree) PreOrder() []interface{} {
	t := bt
	stack := NewStack(reflect.TypeOf(bt))
	res := make([]interface{}, 0)
	for t != nil || !stack.Empty() {
		for t != nil {
			res = append(res, t.Data)
			stack.Push(t)
			t = t.Lchild
		}
		if !stack.Empty() {
			v, _ := stack.Pop()
			t = v.(*BinaryTree)
			t = t.Rchild
		}
	}
	return res
}

// 中序遍历
func (bt *BinaryTree) InOrder() []interface{} {
	t := bt
	stack := NewStack(reflect.TypeOf(bt))
	res := make([]interface{}, 0)
	for t != nil || !stack.Empty() {
		for t != nil {
			stack.Push(t)
			t = t.Lchild
		}
		if !stack.Empty() {
			v, _ := stack.Pop()
			t = v.(*BinaryTree)
			res = append(res, t.Data)
			t = t.Rchild
		}
	}
	return res
}

// 后续遍历
func (bt *BinaryTree) PostOrder() []interface{} {
	t := bt
	stack := NewStack(reflect.TypeOf(bt))
	s := NewStack(reflect.TypeOf(true))
	res := make([]interface{}, 0)
	for t != nil || !stack.Empty() {
		for t != nil {
			stack.Push(t)
			s.Push(false)
			t = t.Lchild
		}
		for flag, _ := s.Top(); !stack.Empty() && flag.(bool); {
			s.Pop()
			v, _ := stack.Pop()
			res = append(res, v.(*BinaryTree).Data)
			flag, _ = s.Top()
		}
		if !stack.Empty() {
			s.Pop()
			s.Push(true)
			v, _ := stack.Top()
			t = v.(*BinaryTree)
			t = t.Rchild
		}
	}
	return res
}

func main() {
	var n3 = &BinaryTree{2, nil, nil}
	var n2 = &BinaryTree{2, nil, nil}
	var t1 = BinaryTree{1, n2, n3}

	var m7 = &BinaryTree{4, nil, nil}
	var m6 = &BinaryTree{3, nil, nil}
	var m5 = &BinaryTree{3, nil, nil}
	var m4 = &BinaryTree{4, nil, nil}
	var m3 = &BinaryTree{2, m6, m7}
	var m2 = &BinaryTree{2, m4, m5}
	var t2 = BinaryTree{1, m2, m3}

	var o2 = &BinaryTree{2, nil, nil}
	var t3 = BinaryTree{1, o2, nil}

	fmt.Println(BinaryTree)
	fmt.Println(BinaryTree.PreOrder())
	fmt.Println(BinaryTree.InOrder())
	fmt.Println(BinaryTree.PostOrder())

}
