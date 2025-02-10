package jianzhi_offer

import "fmt"

// FindValueInMatrix 在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
// 该矩阵从左至右是有序排序的,时间负责度是在 O(max(m,n))
func FindValueInMatrix(data [][]int, dst int) (int, int) {
	rowMax := len(data)
	colMax := len(data[0])
	if rowMax <= 0 || colMax <= 0 {
		return -1, -1
	}

	for i, j := 0, colMax-1; i < rowMax && j >= 0; {
		if data[i][j] == dst {
			fmt.Println("find row: ", i, ", col: ", j)
			return i, j
		}
		if data[i][j] < dst {
			i++
			continue
		}
		j--
	}
	return -1, -1
}
