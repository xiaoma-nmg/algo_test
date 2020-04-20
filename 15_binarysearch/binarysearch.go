package _5_binarysearch

func BinarySearch(data []int, val int) int  {
    left, right := 0, len(data)-1
    for left <= right {
        mid := left + ((right - left) >> 1)
        if data[mid] > val {
            right = mid - 1
        } else if data[mid] < val {
            left = mid + 1
        } else {
            return mid
        }
    }
    return -1
}

func bs(data []int, val int, left, right int) int {
    if left > right {
        return -1
    }
    mid := left + ((right - left) >> 1)
    if data[mid] == val {
        return mid
    }
    if data[mid] < val {
        return bs(data, val, mid + 1, right)
    }
    return bs(data, val, left, mid - 1)
}

func BinarySearchRecursive(data []int, val int) int {
    return bs(data, val, 0, len(data) - 1)
}
