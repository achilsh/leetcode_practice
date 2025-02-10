package jianzhi_offer

// BinarySearch 在有序数组中查找 某个值，存在则返回索引，不存在则返回-1
func BinarySearch(data []int, dst int) int {
	iLight, iRight := 0, len(data)-1

	for {
		if iLight > iRight {
			return -1
		}
		iMiddle := (iLight + iRight) / 2
		if data[iMiddle] == dst {
			return iMiddle
		}
		if data[iMiddle] > dst {
			iRight = iMiddle - 1
			continue
		}
		if data[iMiddle] < dst {
			iLight = iMiddle + 1
		}
	}
}
