package _2_sort

import "testing"

func TestMergeSort(t *testing.T) {
    array := []int{4, 5, 6, 3, 2, 1}
    t.Log(array)
    MergeSort(array)
    t.Log(array)
}


func TestQuickSort(t *testing.T) {
    array := []int{4, 5, 6, 3, 2, 1}
    t.Log(array)
    QuickSort(array)
    t.Log(array)
}
