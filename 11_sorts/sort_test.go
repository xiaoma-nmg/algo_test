package _1_sorts

import "testing"

func TestBubble(t *testing.T) {
	array := []int{4, 5, 6, 3, 2, 1, 99, 97, 46, 32, 55, 66, 31, 78}
	t.Log(array)
	Bubble(array)
	t.Log(array)
}

func TestInsert(t *testing.T) {
	array := []int{4, 5, 6, 3, 2, 1, 99, 97, 46, 32, 55, 66, 31, 78}
	t.Log(array)
	Insert(array)
	t.Log(array)
}

func TestSelect(t *testing.T) {
	array := []int{4, 5, 6, 3, 2, 1, 99, 97, 46, 32, 55, 66, 31, 78}
	t.Log(array)
	Insert(array)
	t.Log(array)
}

func TestShell(t *testing.T) {
	array := []int{4, 5, 6, 3, 2, 1, 99, 97, 46, 32, 55, 66, 31, 78}
	t.Log(array)
	Shell(array)
	t.Log(array)
}
