package _8_stack

import "fmt"

type ArrayStack struct {
    data []interface{}
    top int
}

func NewArrayStack() *ArrayStack {
    return &ArrayStack{
        data: make([]interface{}, 0, 32),
        top:  -1,
    }
}

func (a *ArrayStack)Push(v interface{})  {
    if a.top < 0 {
        a.top = 0
    } else {
        a.top++
    }

    if a.top > len(a.data) - 1 {
        a.data = append(a.data, v)
    } else {
        a.data[a.top] = v
    }
}

func (a *ArrayStack)Pop() interface{} {
    if a.IsEmpty() {
        return nil
    }
    val := a.data[a.top]
    a.top--
    return val
}

func (a *ArrayStack)IsEmpty() bool  {
    if a.top < 0 {
        return true
    }
    return false
}

func (a *ArrayStack)Top() interface{}  {
    if a.IsEmpty() {
        return nil
    }
    return a.data[a.top]
}

func (a *ArrayStack)Flush()  {
    a.top = -1
}

func (a *ArrayStack)Print()  {
    if a.IsEmpty() {
        fmt.Println("empty stack")
        return
    }
    for i:= a.top; i >= 0; i-- {
        fmt.Println(a.data[i])
    }
}
