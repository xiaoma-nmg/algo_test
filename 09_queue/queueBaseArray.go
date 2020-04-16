package _9_queue

import "fmt"

type QueueArray struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

func NewQueueArray(n int) *QueueArray {
	return &QueueArray{
		data:     make([]interface{}, n, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (q *QueueArray) EnQueue(v interface{}) bool {
	if q.tail == q.capacity {
		if q.head == 0 {
			return false
		}
		copy(q.data, q.data[q.head:q.tail])
		q.tail = q.tail - q.head
		q.head = 0
	}
	q.data[q.tail] = v
	q.tail++
	return true
}

func (q *QueueArray) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}
	val := q.data[q.head]
	q.head++
	return val
}

func (q *QueueArray) Print() {
	for i := q.head; i < q.tail; i++ {
		fmt.Printf("%+v\t", q.data[i])
	}
	fmt.Println()
}
