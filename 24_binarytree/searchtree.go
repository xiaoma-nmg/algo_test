package _4_binarytree

import (
    "fmt"
)

type TreeNode struct {
    val int
    left *TreeNode
    right *TreeNode
}

type Tree struct {
    root *TreeNode
}

func NewTreeNode(val int) *TreeNode  {
    return &TreeNode{
        val:   val,
        left:  nil,
        right: nil,
    }
}

func NewTree() *Tree {
    return &Tree{root:nil}
}

func (t *Tree)Insert(val int) {
    if t.root == nil {
        node := NewTreeNode(val)
        t.root = node
    }

    cur := t.root
    for cur != nil {
        if val > cur.val {
            if cur.right == nil {
                node := NewTreeNode(val)
                cur.right = node
                break
            }
            cur = cur.right
        } else if val < cur.val {
            if cur.left == nil {
                node := NewTreeNode(val)
                cur.left = node
                break
            }
            cur = cur.left
        } else {
            break
        }
    }
}

func (t *Tree)Search(val int) bool {
    if t.root == nil {
        return false
    }
    cur := t.root
    for cur != nil {
        if val > cur.val {
            cur = cur.right
        } else if val < cur.val {
            cur = cur.left
        } else {
            return true
        }
    }
    return false
}

func (t *Tree)Delete(val int) {
    if t.root == nil {
        return
    }

    var father *TreeNode = nil
    cur := t.root

    for cur != nil {
        if val > cur.val {
            father = cur
            cur = cur.right
        } else if val < cur.val {
            father = cur
            cur = cur.left
        } else {
            break
        }
    }

    if cur == nil {
        return
    }

    // 如果当前节点没有孩子，父节点指向当前节点的指针置为nil
    if cur.left == nil && cur.right == nil {
        if father == nil {
            t.root = nil
            return
        }

        if father.left != nil && father.left.val == val {
            father.left = nil
            return
        }

        if father.right != nil && father.right.val == val {
            father.right = nil
            return
        }
    }

    // 如果当前节点只有一个孩子，父节点指向当前节点的指针指向当前节点的孩子
    if cur.left == nil && cur.right != nil {
        if father == nil {
            t.root = cur.right
            return
        }
        if father.left.val == val {
            father.left = cur.right
            return
        }

        if father.right.val == val {
            father.right = cur.right
            return
        }
    }

    if cur.left != nil && cur.right == nil {
        if father == nil {
            t.root = cur.left
            return
        }

        if father.left.val == val {
            father.left = cur.left
            return
        }

        if father.right.val == val {
            father.right = cur.left
            return
        }
    }

    // 如果当前结点左右子树都存在，则从右子树中把最小的值，替换过来。再删除当前结点
    rightFather := cur
    right := cur.right
    if right.left != nil {
        rightFather = right
        right = right.left
    }
    cur.val= right.val
    if rightFather == cur {
        cur.right = nil
    } else {
        rightFather.left = nil
    }
}

func print(node *TreeNode)  {
    if node == nil {
        return
    }
    fmt.Printf("%d \t", node.val)
    print(node.left)
    print(node.right)
}

func (t *Tree)Print()  {
    print(t.root)
    fmt.Println()
}