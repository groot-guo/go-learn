package main

// 使用多余得一个矩阵来旋转图像
func rotateImage(matrix [][]int) {
	res := make([][]int, len(matrix))
	for index := range matrix {
		res[index] = append(res[index], matrix[index]...)
	}

	leftIndex := 0
	topIndex := 0
	rightIndex := len(matrix) - 1
	for leftIndex < len(matrix) && topIndex < len(matrix) {
		matrix[topIndex][rightIndex] = res[leftIndex][topIndex]
		topIndex++
		if topIndex == len(matrix) {
			topIndex = 0
			rightIndex--
			leftIndex++
		}
	}

}

// 直接交换
func rotate23(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}
	rows := len(matrix)
	for i := 0; i < rows/2; i++ {
		matrix[i], matrix[rows-1-i] = matrix[rows-1-i], matrix[i]
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return
}
