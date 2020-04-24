package _7_skiplist

import "testing"

func TestSkipList(t *testing.T) {
	s := NewSkipList()
	s.Print()
	for i := 1; i <= 1000; i += 2 {
		s.Insert(i, i)
		//s.Print()
	}
	s.Print()

	if cur, OK := s.Search(556, 555); OK {
		t.Log(cur.val)
	}

	if cur, OK := s.Search(666, 666); OK {
		t.Log(cur.val)
	}

}
