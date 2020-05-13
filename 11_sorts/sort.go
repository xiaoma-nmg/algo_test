package _1_sorts

// 冒泡、插入、选择

func Bubble(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		change := false

		for j := 0; j < length-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				change = true
			}
		}

		if !change {
			break
		}
	}
}

func Insert(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}

	for i := 1; i < length; i++ {
		cur := data[i]
		for j := 0; j < i; j++ {
			if cur < data[j] {
				copy(data[j+1:i+1], data[j:i])
				data[j] = cur
				break
			}
		}
	}
}

func Select(data []int) {
	length := len(data)
	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		min := i + 1
		for j := i + 1; j < length; j++ {
			if data[j] < data[min] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}

func Shell(data []int) {
	length := len(data)
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			for j := i; j >= gap && data[j] < data[j-gap]; j -= gap {
				data[j], data[j-gap] = data[j-gap], data[j]
			}
		}
	}
}
