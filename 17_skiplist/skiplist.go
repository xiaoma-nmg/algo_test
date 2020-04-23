package _7_skiplist

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

const (
	// 最高层数
	MaxLEVEL = 16
)

type SkipListNode struct {
	val      interface{}
	score    int
	level    int
	forwards []*SkipListNode
}

type SkipList struct {
	head   *SkipListNode
	level  int
	length int
}

func NewSkipListNode(val interface{}, score int, level int) *SkipListNode {
	return &SkipListNode{
		val:      val,
		score:    score,
		level:    level,
		forwards: make([]*SkipListNode, level),
	}
}

func NewSkipList() *SkipList {
	head := NewSkipListNode(0, math.MinInt32, MaxLEVEL)
	return &SkipList{
		head:   head,
		level:  1,
		length: 0,
	}
}

func (s *SkipList) Level() int {
	return s.level
}

func (s *SkipList) Length() int {
	return s.length
}

func (s *SkipList) Insert(val interface{}, score int) error {
	if val == nil {
		return errors.New("can not insert nil value")
	}

	cur := s.head
	levelStart := [MaxLEVEL]*SkipListNode{}

	for i := MaxLEVEL - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].val == val {
				return errors.New("value exist")
			}
			if cur.forwards[i].score > score {
				levelStart[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			levelStart[i] = cur
		}
	}

	// 通过随机算法获取该节点的层数
	level := 1
	for i := 1; i < MaxLEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	newNode := NewSkipListNode(val, score, level)
	//  插入
	for i := 0; i <= level-1; i++ {
		next := levelStart[i].forwards[i]
		levelStart[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	// 更新整个表的level 以及 length
	if level > s.level {
		s.level = level
	}

	s.length++
	return nil
}

func (s *SkipList) Search(val interface{}, score int) *SkipListNode {
	if val == nil || s.length == 0 {
		return nil
	}

	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].val == val {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

func (s *SkipList) Delete(val interface{}, score int) error {
	if val == nil {
		return errors.New("value is nil")
	}

	cur := s.head
	levelStart := [MaxLEVEL]*SkipListNode{}
	for i := s.level - 1; i >= 0; i-- {
		levelStart[i] = s.head
		for cur.forwards[i] != nil {
			if cur.forwards[i].score == score && cur.forwards[i].val == val {
				levelStart[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = levelStart[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if levelStart[i] == s.head && cur.forwards[i] == nil {
			s.level = i
		}

		if levelStart[i].forwards[i] != nil {
			levelStart[i].forwards[i] = levelStart[i].forwards[i].forwards[i]
		}
	}
	s.length--
	return nil
}

func (s *SkipList) Print() {
	fmt.Printf("skip list has %d level and length is %d\n", s.level, s.length)
	for i := s.level - 1; i >= 0; i-- {
		cur := s.head.forwards[i]
		for cur != nil {
			fmt.Printf("%#v\t", cur.val)
			cur = cur.forwards[i]
		}
		fmt.Println("\n---------------------------------------------------------")
	}
}
