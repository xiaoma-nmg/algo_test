package _2_sort

import "testing"

func TestFindK(t *testing.T) {
    array := []int{11, 8, 3, 9, 1, 7, 2, 5}

    for i:=1; i<= len(array); i++ {
        t.Log(FindK(array, i))
    }
}
