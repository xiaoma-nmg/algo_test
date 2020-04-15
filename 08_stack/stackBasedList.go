package _8_stack

import "fmt"

// 基于链表实现栈

type node struct {
    next *node
    val interface{}
}

type LinkedListStack struct {
    top *node
}

func NewLinkStack() *LinkedListStack {
    return &LinkedListStack{nil}
}

func (l *LinkedListStack)Push(v interface{})  {
    l.top = &node{
        next: l.top,
        val:  v,
    }
}

func (l *LinkedListStack)Pop() interface{} {
    if l.IsEmpty() {
        return nil
    }
    val := l.top.val
    l.top = l.top.next
    return val
}

func (l *LinkedListStack) IsEmpty() bool {
    return l.top == nil
}

func (l *LinkedListStack) Top() interface{} {
    if l.IsEmpty() {
        return nil
    }
    return l.top.val
}

func (l *LinkedListStack) Flush()  {
    l.top = nil
}

func (l *LinkedListStack) Print()  {
    if l.IsEmpty() {
        fmt.Println("empty stack")
        return
    }

    cur := l.top
    for cur != nil {
        fmt.Println(cur.val)
        cur = cur.next
    }
}