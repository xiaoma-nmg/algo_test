package _2_sort

func FindK(data []int, k int) int {
    if k > len(data) {
        return 0
    }
    return find(data, 0, len(data) - 1, k)
}

func find(data[]int, start, end, k int) int {
    if start >= end {
        return data[end]
    }

    i := start
    for j:=start; j<end; j++ {
        if data[j] > data[end]{
            data[j],data[i] = data[i], data[j]
            i++
        }
    }
    data[i], data[end] = data[end], data[i]

    if i + 1 > k {
        return find(data, start, i-1, k)
    } else if i + 1 < k {
        return find(data, i+1, end, k)
    } else {
        return data[i]
    }
}
