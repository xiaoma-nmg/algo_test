package _9_queue

import "fmt"

type QueueCircular struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

func NewQueueCircular(n int) *QueueCircular {
	return &QueueCircular{
		data:     make([]interface{}, n, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (q *QueueCircular) IsEmpty() bool {
	return q.head == q.tail
}

func (q *QueueCircular) IsFull() bool {
	return (q.tail+1)%q.capacity == q.head
}

func (q *QueueCircular) EnQueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *QueueCircular) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	v := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	return v
}

func (q *QueueCircular) Print() {
	if q.IsEmpty() {
		fmt.Println("empty queue")
		return
	}

	for i := q.head; i != q.tail; i = (i + 1) % q.capacity {
		fmt.Printf("%+v\t", q.data[i])
	}
	fmt.Println()

}
