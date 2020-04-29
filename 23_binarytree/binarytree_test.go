package _3_binarytree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := CreateTree()
	root.PreOrder(root)
	fmt.Println()
	root.InOrder(root)
	fmt.Println()
	root.PostOrder(root)
	fmt.Println()
	root.LevelOrder(root)
	fmt.Println()
}
