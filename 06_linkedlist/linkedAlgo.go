package _6_linkedlist

/**
 * 1) 单链表反转
 * 2) 链表中环的检测
 * 3) 两个有序的链表合并
 * 4) 删除链表倒数第n个结点
 * 5) 求链表的中间结点
 *
 */


func (l *LinkedList) Reverse() {
    var pre *ListNode = nil
    cur := l.head.Next

    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    l.head.Next = pre
}

func (l *LinkedList)HasLoop() bool {
    if l.head.Next == nil {
        return false
    }
    s, q := l.head.Next, l.head.Next.Next
    for q != nil && q.Next != nil {
        if q == s {
            return true
        }
        s = s.Next
        q = q.Next.Next
    }
    return false
}

func MergeOrderList(l1 *LinkedList, l2 *LinkedList) *LinkedList {
    p1 := l1.head.Next
    p2 := l2.head.Next
    l := NewLinkedList()
    cur := l.head

    for p1 != nil && p2 != nil {
        if p1.GetValue().(int) < p2.GetValue().(int) {
            cur.Next = p1
            p1 = p1.Next
        } else {
            cur.Next = p2
            p2 = p2.Next
        }
        cur = cur.Next
    }
    if p1 != nil {
        cur.Next = p1
    } else if p2 != nil {
        cur.Next = p2
    }
    return l
}

func (l *LinkedList)DeleteTailN(n int)  {
    s, q := l.head.Next, l.head.Next
    for n > 0 {
        q = q.Next
        n--
    }

    for q.Next != nil {
        s = s.Next
        q = q.Next
    }
    s.Next = s.Next.Next
}

func (l *LinkedList)GetMidNode() interface{} {
    s, q := l.head, l.head.Next
    for q != nil && q.Next != nil {
        q = q.Next.Next
        s = s.Next
    }
    return s.GetValue()
}