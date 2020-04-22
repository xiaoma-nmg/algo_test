package _6_binarysearch

import "testing"

var data []int = []int{1,3,4,5,6,8,8,8,11,18}

func TestFindFirstKey(t *testing.T) {
    t.Log(data)
    t.Log(FindFirstKey(data, 8))
}

func TestFindLastKey(t *testing.T) {
    t.Log(data)
    t.Log(FindLastKey(data, 8))
    t.Log(FindLastKey(data, 3))
    t.Log(FindLastKey(data, 18))
}

func TestFindFirstLargeVal(t *testing.T) {
    t.Log(data)
    t.Log(FindFirstLargeVal(data, 5))
    t.Log(FindFirstLargeVal(data, 8))
    t.Log(FindFirstLargeVal(data, 18))
    t.Log(FindFirstLargeVal(data, 2))
}

func TestFindLastLessVal(t *testing.T) {
    t.Log(data)
    t.Log(FindLastLessVal(data, 5))
    t.Log(FindLastLessVal(data, 8))
    t.Log(FindLastLessVal(data, 18))
    t.Log(FindLastLessVal(data, 2))
}