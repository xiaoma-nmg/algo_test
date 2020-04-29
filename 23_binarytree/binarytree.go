package _3_binarytree

import "fmt"

type Tree struct {
	val   interface{}
	left  *Tree
	right *Tree
}

func NewTreeNode(val interface{}) *Tree {
	return &Tree{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (t *Tree) PreOrder(root *Tree) {
	if root == nil {
		return
	}
	fmt.Printf("%v\t", root.val)
	t.PreOrder(root.left)
	t.PreOrder(root.right)
}

func (t *Tree) InOrder(root *Tree) {
	if root == nil {
		return
	}
	t.InOrder(root.left)
	fmt.Printf("%v\t", root.val)
	t.InOrder(root.right)
}

func (t *Tree) PostOrder(root *Tree) {
	if root == nil {
		return
	}
	t.PostOrder(root.left)
	t.PostOrder(root.right)
	fmt.Printf("%v\t", root.val)
}

func (t *Tree) LevelOrder(root *Tree) {
	if root == nil {
		return
	}
	queue := make([]*Tree, 0, 32)
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Printf("%v\t", cur.val)
		if cur.left != nil {
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			queue = append(queue, cur.right)
		}
	}
}

func CreateTree() *Tree {
	root := NewTreeNode(1)
	queue := make([]*Tree, 0, 32)
	queue = append(queue, root)
	for i := 2; i < 10; i++ {
		cur := queue[0]
		queue = queue[1:]
		cur.left = NewTreeNode(i)
		i++
		cur.right = NewTreeNode(i)
		queue = append(queue, cur.left, cur.right)
	}
	return root
}
