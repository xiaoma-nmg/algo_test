package _6_linkedlist

import "fmt"

func IsPalindrome(l *LinkedList) bool {
    if l.length == 0 {
        return false
    }
    if l.length == 1 {
        return true
    }

    // 找到链表中间结点
    q, s := l.head.Next.Next, l.head.Next
    for q != nil && q.Next != nil {
        s = s.Next
        q = q.Next.Next
    }

    // 翻转链表后半段
    var pre *ListNode = nil
    cur := s.Next
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    fmt.Println("翻转完毕", "尾端头部：", pre.val, "s 是", s.val)

    head := l.head.Next
    for pre != nil {
        if pre.val != head.val {
            return false
        }
        pre = pre.Next
        head = head.Next
    }
    return true
}