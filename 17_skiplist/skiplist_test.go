package _7_skiplist

import "testing"

func TestSkipList(t *testing.T) {
	s := NewSkipList()
	s.Print()
	for i := 1; i <= 100; i++ {
		s.Insert(i, i)
		//s.Print()
	}
	s.Print()
}
