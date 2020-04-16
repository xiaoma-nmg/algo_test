package _9_queue

import "testing"

func TestQueueList_EnQueue(t *testing.T) {
	queue := NewQueueList()
	queue.Print()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.Print()
}

func TestQueueList_DeQueue(t *testing.T) {
	queue := NewQueueList()
	queue.Print()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.Print()
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	queue.Print()
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	queue.Print()
}
