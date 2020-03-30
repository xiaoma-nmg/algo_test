package _6_linkedlist

import "fmt"

type ListNode struct {
    Next *ListNode
    val interface{}
}

type LinkedList struct {
    head *ListNode
    length uint
}


func NewListNode(v interface{}) *ListNode {
    return &ListNode{
        Next: nil,
        val:  v,
    }
}

func (l *ListNode)GetNext() *ListNode {
    return l.Next
}

func (l *ListNode)GetValue() interface{} {
    return l.val
}

func NewLinkedList() *LinkedList {
    return &LinkedList{
        head:   NewListNode(0),
        length: 0,
    }
}

func (l *LinkedList)InsertAfter(node *ListNode, val interface{}) bool {
    if node == nil {
        return false
    }
    newNode := NewListNode(val)
    newNode.Next = node.Next
    node.Next = newNode
    l.length++
    return true
}

func (l *LinkedList)InsertBefore(node *ListNode, val interface{}) bool {
    if node == nil || node == l.head {
        return false
    }

    pre, cur := l.head, l.head.Next
    for cur != nil {
        if cur == node {
            newNode := NewListNode(val)
            pre.Next = newNode
            newNode.Next = cur
            l.length++
            return true
        }
        pre = cur
        cur = cur.Next
    }
    return false
}

func (l *LinkedList)InsertToHead(val interface{}) bool {
    return l.InsertAfter(l.head, val)
}

func (l *LinkedList)GetLinedListTailNode() *ListNode {
    cur := l.head
    for cur.Next != nil {
        cur = cur.Next
    }
    return cur
}

func (l *LinkedList)InsertToTail(val interface{}) bool {
    return l.InsertAfter(l.GetLinedListTailNode(), val)
}

func (l *LinkedList)FindByIndex(index uint) *ListNode {
    if index >= l.length {
        return nil
    }
    cur := l.head.Next
    var i uint = 0
    for ; i<index; i++ {
        cur = cur.Next
    }
    return cur
}

func (l *LinkedList)DeleteNode(node *ListNode) bool {
    if node == nil {
        return false
    }
    pre, cur := l.head, l.head.Next
    for cur != nil {
        if cur == node {
            pre.Next = cur.Next
            cur = nil
            l.length--
            return true
        }
        pre = cur
        cur = cur.Next
    }
    return false
}

func (l *LinkedList)Print() {
    cur := l.head.Next
    format := ""
    for cur != nil {
        format += fmt.Sprintf("%+v->", cur.GetValue())
        cur = cur.Next
    }
    fmt.Println(format[:len(format)-2])
}