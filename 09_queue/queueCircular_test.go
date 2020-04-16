package _9_queue

import "testing"

func TestQueueCircular_EnQueue(t *testing.T) {
	q := NewQueueCircular(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.Print()
}

func TestQueueCircular_DeQueue(t *testing.T) {
	q := NewQueueCircular(5)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	q.Print()
	t.Log(q.DeQueue())
	t.Log(q.DeQueue())
	q.Print()
	t.Log(q.DeQueue())
	q.EnQueue(5)
	q.EnQueue(6)
	q.Print()
}
