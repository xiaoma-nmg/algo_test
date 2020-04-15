package _8_stack

import "testing"

func TestArrayStack_Push(t *testing.T) {
    stack := NewArrayStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    stack.Push(4)
    stack.Print()
}

func TestArrayStack_Pop(t *testing.T) {
    stack := NewArrayStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    stack.Push(4)
    stack.Print()
    t.Log(stack.Pop())
    t.Log(stack.Pop())
    t.Log(stack.Pop())
    t.Log(stack.Pop())
    stack.Print()
}

func TestArrayStack_Top(t *testing.T) {
    stack := NewArrayStack()
    stack.Push(1)
    t.Log(stack.Top())
    stack.Push(2)
    t.Log(stack.Top())
    stack.Push(3)
    t.Log(stack.Top())
    stack.Push(4)
    t.Log(stack.Top())
    stack.Print()
}