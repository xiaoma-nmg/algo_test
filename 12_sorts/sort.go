package _2_sort

//归并排序， 快速排序

func merge(data []int, x1, y1, x2, y2 int)  {
    array := make([]int, y2-x1+1)
    indexArray := 0
    indexX1 := x1
    indexX2 := x2
    for indexX1 <= y1 && indexX2 <= y2 {
        if data[indexX1] < data[indexX2] {
            array[indexArray] = data[indexX1]
            indexX1++
        } else {
            array[indexArray] = data[indexX2]
            indexX2++
        }
        indexArray++
    }

    if indexX1 <= y1 {
        copy(array[indexArray:], data[indexX1:y1+1])
    }

    if indexX2 <= y2 {
        copy(array[indexArray:], data[indexX2:y2+1])
    }

    copy(data[x1:], array)
}

func splitData(data []int, start, end int)  {
    if start >= end {
        return
    }
    mid := start + (end - start)/2
    splitData(data, start, mid)
    splitData(data, mid + 1, end)
    merge(data, start, mid, mid+1, end)
}

func MergeSort(data []int)  {
    splitData(data, 0, len(data) - 1)
}

func quick(data []int, start, end int)  {
    if start >= end {
        return
    }
    i := start
    for j := start; j < end; j++ {
        if data[j] > data[end] {
            continue
        }
        data[i], data[j] = data[j], data[i]
        i++
    }
    data[i], data[end] = data[end], data[i]
    quick(data, start, i-1)
    quick(data, i+1, end)
}

func QuickSort(data []int)  {
    quick(data, 0, len(data) - 1)
}