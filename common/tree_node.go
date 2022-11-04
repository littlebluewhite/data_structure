package common

import "fmt"

type treeNode[T comparable] struct {
	value T
	left  *treeNode[T]
	right *treeNode[T]
}

type TreeNode[T comparable] interface {
	SetLeft(TreeNode[T])
	SetRight(TreeNode[T])
	BreadthFirstPrint()
	BreadthFirstSearch(v T) bool
	DepthFirstPrint()
	PreOrder()
	InOrder()
	PostOrder()
}

func NewTreeNode[T comparable](value T) (t TreeNode[T]) {
	t = &treeNode[T]{value: value}
	return
}

func (t *treeNode[T]) SetLeft(t2 TreeNode[T]) {
	t3 := t2.(*treeNode[T])
	t.left = t3
}

func (t *treeNode[T]) SetRight(t2 TreeNode[T]) {
	t3 := t2.(*treeNode[T])
	t.right = t3
}

func (t *treeNode[T]) BreadthFirstPrint() {
	q := []*treeNode[T]{t}
	for len(q) > 0 {
		pick := q[0]
		q = q[1:]
		fmt.Println(pick.value)
		if pick.left != nil {
			q = append(q, pick.left)
		}
		if pick.right != nil {
			q = append(q, pick.right)
		}
	}
}

func (t *treeNode[T]) BreadthFirstSearch(value T) bool {
	q := []*treeNode[T]{t}
	for len(q) > 0 {
		pick := q[0]
		if pick.value == value {
			return true
		}
		q = q[1:]
		if pick.left != nil {
			q = append(q, pick.left)
		}
		if pick.right != nil {
			q = append(q, pick.right)
		}
	}
	return false
}

// DepthFirstPrint DepthFirstPrint1 time: O(n) space: O(n)
func (t *treeNode[T]) DepthFirstPrint() {
	stack := []*treeNode[T]{t}
	for len(stack) > 0 {
		l := len(stack)
		pick := stack[l-1]
		stack = stack[:l-1]
		fmt.Println(pick.value)
		if pick.right != nil {
			stack = append(stack, pick.right)
		}
		if pick.left != nil {
			stack = append(stack, pick.left)
		}
	}
}

// DepthFirstPrint2 time: O(n) space: O(n)
func DepthFirstPrint2[T comparable](t TreeNode[T]) {
	t2 := t.(*treeNode[T])
	if t2 == nil {
		return
	}
	fmt.Println(t2.value)
	DepthFirstPrint2[T](t2.left)
	DepthFirstPrint2[T](t2.right)
}

func (t *treeNode[T]) PreOrder() {
	fmt.Println(t.value)
	if t.left != nil {
		t.left.PreOrder()
	}
	if t.right != nil {
		t.right.PreOrder()
	}
}

func (t *treeNode[T]) PostOrder() {
	if t.left != nil {
		t.left.PostOrder()
	}
	fmt.Println(t.value)
	if t.right != nil {
		t.right.PostOrder()
	}
}

func (t *treeNode[T]) InOrder() {
	if t.left != nil {
		t.left.InOrder()
	}
	if t.right != nil {
		t.right.InOrder()
	}
	fmt.Println(t.value)
}
