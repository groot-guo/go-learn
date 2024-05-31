package main

import "strconv"

/*
请你判断一个9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字1-9在每一行只能出现一次。
数字1-9在每一列只能出现一次。
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）

*/

func main() {

}

// 无法判断 重复得，
func isValidSudoku2(board [][]byte) bool {
	topIndex := 0
	leftIndex := 1
	keyMap := make(map[byte]int)
	for leftIndex < 9 && topIndex < 9 {
		if string(board[leftIndex][topIndex]) != "." {
			keyMap[board[leftIndex][topIndex]] += 1
		}

		topIndex++
		if topIndex == 9 {
			topIndex = 0
			leftIndex++
		}
	}

	for _, value := range keyMap {
		if value > 9 {
			return false
		}
	}

	return true
}

// 用三个数组来判断， 未解决
func isValidSudoku(board [][]byte) bool {
	topIndex := 0
	leftIndex := 0
	line := make([][]bool, len(board))
	column := make([][]bool, len(board))
	cell := make([][]bool, len(board))
	for leftIndex < 9 && topIndex < 9 {
		str := string(board[leftIndex][topIndex])
		num, _ := strconv.Atoi(str)
		num -= 1
		k := leftIndex/3*3 + topIndex/3
		if str != "." {
			if cell[k] != nil {
				if cell[k][num] || column[topIndex][num] || line[leftIndex][num] {
					return false
				}
			}

			data := cell[k]
			if cell[k] == nil {
				data = make([]bool, 9)
				data[num] = true
			} else {
				data[num] = true
			}
			cell[k] = data
			column[topIndex] = data
			line[leftIndex] = data
		}

		topIndex++
		if topIndex == 9 {
			topIndex = 0
			leftIndex++
		}
	}

	return true
}

func isValidSudoku3(board [][]byte) bool {
	mapTriple := make([]map[byte]string, 3)
	for i := 0; i < len(board); i++ {
		mapRow := make(map[byte]string)
		mapCol := make(map[byte]string)
		if i%3 == 0 {
			//删除之前mapTriple中的元素，重新申请内存空间
			//mapTriple=make([]map[byte]string,3)
			for k := 0; k < 3; k++ {
				mapTriple[k] = make(map[byte]string)
			}
		}
		for j := 0; j < len(board[0]); j++ {
			//行合法
			if board[i][j] != '.' {
				if _, ok := mapRow[board[i][j]]; ok {
					return false
				} else {
					mapRow[board[i][j]] = ""
				}
			}
			//列合法
			if board[j][i] != '.' {
				if _, ok := mapCol[board[j][i]]; ok {
					return false
				} else {
					mapCol[board[j][i]] = ""
				}
			}
			//3X3方阵合法
			if board[i][j] != '.' {
				if _, ok := mapTriple[j/3][board[i][j]]; ok {
					return false
				} else {
					mapTriple[j/3][board[i][j]] = ""
				}
			}
		}
	}
	return true
}
