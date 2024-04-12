package mid_suanfa

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, 0
	for i < len(matrix) {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			break
		}
		i++
	}
	if i == len(matrix) {
		i = len(matrix) - 1
	}

	for i >= 0 {
		for j < len(matrix[0]) {
			if matrix[i][j] == target {
				return true
			}

			if matrix[i][j] > target {
				break
			}
			j++
		}
		i--
		j = 0
	}

	return false
}

func searchMatrix55(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, v := range row {
			if v == target {
				return true
			}
		}
	}
	return false
}

// z 字形查找
func searchMatrix66(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

func searchMatrix662(matrix [][]int, target int) bool {
	for _, row := range matrix {
		i := sort.SearchInts(row, target)
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}
