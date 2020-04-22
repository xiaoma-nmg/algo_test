package _6_binarysearch

/*
    1。查找第一个值等于给定值的元素
    2。查找最后一个值等于给定值的元素
    3。查找第一个大于等于给定值的元素
    4。查找最后一个小于等于给定值的元素
*/

func FindFirstKey(data []int, val int) int {
    left, right := 0, len(data)-1
    for left <= right {
        mid := left + ((right - left) >> 1)
        if data[mid] > val {
            right = mid - 1
        } else if data[mid] < val {
            left = mid + 1
        } else {
            // naive method
            //for mid - 1 > 0 && data[mid - 1] == val {
            //    mid--
            //}
            //return mid

            if mid == 0 || data[mid - 1] != val {
                return mid
            }
            right = mid - 1
        }
    }
    return -1
}

func FindLastKey(data []int, val int) int  {
    left, right := 0, len(data) - 1
    for left < right {
        mid := left + ((right - left) >> 1)
        if data[mid] > val {
            right = mid - 1
        } else if data[mid] < val {
            left = mid + 1
        } else {
            if mid == len(data) - 1 || data[mid + 1] != val {
                return mid
            }
            left = mid + 1
        }
    }
    return -1
}

func FindFirstLargeVal(data []int, val int) int {
    left, right := 0, len(data) - 1
    for left <= right {
        mid := left + ((right - left) >> 1)
        if data[mid] >= val {
            if mid == 0 || data[mid - 1] < val {
                return mid
            }
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}

func FindLastLessVal(data []int, val int) int {
    left, right := 0, len(data) - 1
    for left <= right {
        mid := left + ((right - left) >> 1)
        if data[mid] > val {
            right = mid - 1
        } else {
            if mid == len(data) - 1 || data[mid + 1] > val {
                return mid
            }
            left = mid + 1
        }
    }
    return -1
}