package _9_queue

import "testing"

func TestQueueArray_EnQueue(t *testing.T) {
	q := NewQueueArray(8)
	q.EnQueue(1)
	q.Print()
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.EnQueue(7)
	q.EnQueue(8)
	q.Print()
	q.EnQueue(9)
	q.Print()
}

func TestQueueArray_DeQueue(t *testing.T) {
	q := NewQueueArray(8)
	q.EnQueue(1)
	q.Print()
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.EnQueue(7)
	q.EnQueue(8)
	q.Print()
	q.EnQueue(9)
	q.Print()
	t.Log(q.DeQueue())
	t.Log(q.DeQueue())
	t.Log(q.DeQueue())
	q.Print()
	q.EnQueue(9)
	q.EnQueue(10)
	q.Print()
	q.EnQueue(11)
	q.Print()
	q.EnQueue(12)
	q.Print()
}
