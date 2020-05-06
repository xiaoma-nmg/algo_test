package _8_heap

import "testing"

func TestHeap_Insert(t *testing.T) {
	heap := NewHeap(10)
	heap.Insert(5)
	heap.Insert(10)
	heap.Insert(8)
	heap.Print()
	heap.Insert(9)
	heap.Print()
	heap.Insert(4)
	heap.Print()
	heap.Delete()
	heap.Print()
	heap.Delete()
	heap.Print()
	heap.Delete()
	heap.Print()
}
