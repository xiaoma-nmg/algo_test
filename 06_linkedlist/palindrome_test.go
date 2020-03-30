package _6_linkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
    str := []string{"12345", "123321", "heoeh", "a", "12321"}
    for _, s := range str {
        l := NewLinkedList()
        for _, c := range s {
            l.InsertToTail(string(c))
        }
        l.Print()
        t.Log(IsPalindrome(l))
        l.Print()
    }
}