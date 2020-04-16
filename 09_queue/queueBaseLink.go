package _9_queue

import "fmt"

type ListNode struct {
	next *ListNode
	val  interface{}
}

type QueueList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewQueueList() *QueueList {
	return &QueueList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *QueueList) EnQueue(v interface{}) {
	node := &ListNode{
		next: nil,
		val:  v,
	}

	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

func (q *QueueList) DeQueue() interface{} {
	if q.head == nil {
		return nil
	}
	val := q.head.val
	q.head = q.head.next
	q.length--
	return val
}

func (q *QueueList) Print() {
	cur := q.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

func (q *QueueList) GetLength() int {
	return q.length
}
