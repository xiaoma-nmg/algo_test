package _8_heap

import (
	"errors"
	"fmt"
)

type Heap struct {
	data     []int
	length   int
	capacity int
}

func NewHeap(capacity int) *Heap {
	return &Heap{
		data:     make([]int, capacity, capacity),
		length:   0,
		capacity: capacity,
	}
}

func (h *Heap) Insert(val int) error {
	if h.length == h.capacity {
		return errors.New("heap is full")
	}

	// 把数据放进数组末尾
	h.data[h.length] = val
	h.length++

	// 堆化
	cur := h.length - 1
	father := (cur - 1) / 2
	for father >= 0 && h.data[cur] > h.data[father] {
		h.data[cur], h.data[father] = h.data[father], h.data[cur]
		cur = father
		father = (cur - 1) / 2
	}
	return nil
}

func (h *Heap) Delete() error {
	if h.length <= 0 {
		return errors.New("heap is empty")
	}
	// 将数组最后一个元素放到堆顶
	h.length--
	h.data[0] = h.data[h.length]

	// 堆化
	cur := 0
	for {
		// 当前节点以及它的左右孩子中选出最大的值
		maxPos := cur
		if 2*cur+1 < h.length && h.data[maxPos] < h.data[2*cur+1] {
			maxPos = 2*cur + 1
		}
		if 2*cur+2 < h.length && h.data[maxPos] < h.data[2*cur+2] {
			maxPos = 2*cur + 2
		}

		// 当前节点比左右子树的节点大，堆化完成
		if maxPos == cur {
			break
		}

		h.data[cur], h.data[maxPos] = h.data[maxPos], h.data[cur]
		cur = maxPos
	}

	return nil
}

func (h *Heap) Print() {
	for i := 0; i < h.length; i++ {
		fmt.Printf("%d \t", h.data[i])
	}
	fmt.Println()
}
