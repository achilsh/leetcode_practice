package jianzhi_offer

import "fmt"

// 找出数组红重复的数字？
// 在一个长度为 n 的数组里的所有数字都在 0～n - 1 的范围内。数组中某些数字是重复的，
// 但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
// 例如，如果输入长度为 7 的数组 {2, 3, 1, 0, 2, 5, 3}，那么对应的输出是重复的数字 2 或者 3。

// FindRepeatedNumInArray 时间复杂度是 O(n)
func FindRepeatedNumInArray(data []int) int {
	srcArray := data[:]
	if len(srcArray) <= 0 {
		return -1
	}

	for i := 0; i < len(srcArray); i++ {
		if srcArray[i] < 0 || srcArray[i] > len(srcArray)-1 {
			return -1
		}
	}
	for i := 0; i < len(srcArray); i++ {
		for {
			//判断数组下标值和数组值是否相等，不相等则把数组值，把数值和下标为数组值的数组值互换一下。
			// 相等则数组下一个元素进行检查
			if srcArray[i] != i { //1 != 0
				dataI := srcArray[i]      // 2// 1
				dataII := srcArray[dataI] //1// 3

				if dataII == dataI {
					fmt.Println("find same value: ", dataI)
					return dataI
				}
				srcArray[i] = dataII // 1 //
				srcArray[dataI] = dataI
			} else {
				break
			}
		}
	}
	//var i = 0
	//for {
	//	if srcArray[i] != i { //1 != 0
	//		dataI := srcArray[i]      // 2// 1
	//		dataII := srcArray[dataI] //1// 3
	//
	//		if dataII == dataI {
	//			fmt.Println("find same value: ", dataI)
	//			return dataI
	//		}
	//		srcArray[i] = dataII // 1 //
	//		srcArray[dataI] = dataI
	//		continue
	//	}
	//	i++
	//	if i >= len(srcArray) {
	//		break
	//	}
	//}
	return -1
}
