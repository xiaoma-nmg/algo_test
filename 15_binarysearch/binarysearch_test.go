package _5_binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
    data := []int{1,2,3,4,5,6,7,8,9}
    t.Log(BinarySearch(data, 3))
    t.Log(BinarySearch(data, 0))
    t.Log(BinarySearch(data, 10))
    t.Log(BinarySearchRecursive(data, 5))
    t.Log(BinarySearchRecursive(data, 0))
    t.Log(BinarySearchRecursive(data, 10))
}
