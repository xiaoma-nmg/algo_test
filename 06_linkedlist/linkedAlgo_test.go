package _6_linkedlist

import "testing"

func TestReverse(t *testing.T) {
    l := NewLinkedList()
    for i := 0; i < 10; i++ {
        l.InsertToTail(i + 1)
    }
    l.Print()
    l.Reverse()
    l.Print()
}

func TestHasLoop(t *testing.T) {
    l := NewLinkedList()
    for i := 0; i < 10; i++ {
        l.InsertToTail(i + 1)
    }
    l.Print()
    t.Log(l.HasLoop())
    pre := l.head.Next.Next.Next.Next
    t.Log(pre.GetValue())
    tail := l.head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next
    t.Log(tail.GetValue())
    tail.Next = pre
    t.Log(l.HasLoop())
}

func TestMergeOrderList(t *testing.T) {
    l1 := NewLinkedList()
    l2 := NewLinkedList()
    for i:=1; i<20; i = i+2 {
        l1.InsertToTail(i)
        l2.InsertToTail(i+1)
    }
    l1.Print()
    l2.Print()
    tmp := MergeOrderList(l1, l2)
    tmp.Print()
}

func TestDeleteTailN(t *testing.T)  {
    l := NewLinkedList()
    for i:=1; i<10; i++ {
        l.InsertToTail(i)
    }
    l.Print()
    l.DeleteTailN(2)
    l.Print()
    l.DeleteTailN(4)
    l.Print()
}

func TestGetMidNode(t *testing.T) {
    l := NewLinkedList()
    for i:=1; i<9; i++ {
        l.InsertToTail(i)
    }
    l.Print()
    t.Log(l.GetMidNode())
}