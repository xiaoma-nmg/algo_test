package _8_stack

import "fmt"

type Browser struct {
    forwardStack Stack
    backwardStack Stack
}

func NewBrowser() *Browser {
    return &Browser{
        forwardStack:  NewArrayStack(),
        backwardStack: NewLinkStack(),
    }
}

func (b *Browser)CanForward() bool {
    if b.forwardStack.IsEmpty() {
        return false
    }
    return true
}

func (b *Browser)CanBackward() bool {
    if b.backwardStack.IsEmpty() {
        return false
    }
    return true
}

func (b *Browser) Open(addr string)  {
    fmt.Printf("Open new addr %+v\n", addr)
    b.forwardStack.Flush()
}

func (b *Browser) PushBack(addr string)  {
    b.backwardStack.Push(addr)
}

func (b *Browser) Forward()  {
    if !b.CanForward() {
        return
    }
    top := b.forwardStack.Pop()
    b.backwardStack.Push(top)
    fmt.Printf("forward to %+v\n", top)
}

func (b *Browser)Back()  {
    if !b.CanBackward() {
        return
    }
    top := b.backwardStack.Pop()
    b.forwardStack.Push(top)
    fmt.Printf("backward to %+v\n", top)
}

