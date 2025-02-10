package jianzhi_offer

// BubbleSort 冒泡排序，时间复杂度是 0(n2)
func BubbleSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		for j := 0; j < len(data)-i; j++ {
			if data[j] > data[j+1] {
				tmpData := data[j]
				data[j] = data[j+1]
				data[j+1] = tmpData
			}
		}
	}
	return data
}

// InsertSort 插入排序，时间复杂度是 O(n2)
func InsertSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		var toInsertValue = data[i]
		var j = i - 1
		for ; j >= 0; j-- {
			if data[j] > toInsertValue {
				data[j+1] = data[j]
				continue
			}
			data[j+1] = toInsertValue
			break
		}
		if j <= 0 {
			data[j+1] = toInsertValue
		}
	}
	return data
}

// ChoseSort 选择排序法
func ChoseSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		var maxIndex = 0
		for j := 0; j <= len(data)-i; j++ {
			if data[maxIndex] < data[j] {
				maxIndex = j
			}
		}
		if maxIndex != len(data)-i {
			tmpData := data[maxIndex]
			data[maxIndex] = data[len(data)-i]
			data[len(data)-i] = tmpData
		}

	}
	return data
}
