package _4_binarytree

import (
    "testing"
)

func TestTree_Insert(t *testing.T) {
    tree := NewTree()
    tree.Insert(33)
    tree.Insert(16)
    tree.Insert(50)
    tree.Insert(13)
    tree.Insert(18)
    tree.Insert(15)
    tree.Insert(17)
    tree.Insert(25)
    tree.Insert(19)
    tree.Insert(27)
    tree.Insert(34)
    tree.Insert(58)
    tree.Insert(51)
    tree.Insert(55)
    tree.Insert(66)
    tree.Print()


    tree.Delete(13)
    tree.Print()
    tree.Delete(18)
    tree.Print()
    tree.Delete(55)
    tree.Print()
    tree.Delete(33)
    tree.Print()
}
