package mid_suanfa

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 运行效率最快
func exist22(board [][]byte, word string) bool {
	row, col, wLen := len(board), len(board[0]), len(word)
	target := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	if row*col < wLen {
		return false
	}
	wMap := make(map[byte]int)
	for i := 0; i < wLen; i++ {
		wMap[word[i]]++
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			wMap[board[i][j]]--
		}
	}
	for _, v := range wMap {
		if v > 0 {
			return false
		}
	}
	flag := make([][]bool, row)
	for i := 0; i < row; i++ {
		flag[i] = make([]bool, col)
	}
	var aa func(i, j, l int) bool
	aa = func(i, j, l int) bool {
		if l == len(word) {
			return true
		}
		f := false
		for k := 0; k < 4; k++ {
			tmpI, tmpJ := i+target[k][0], j+target[k][1]
			if (tmpI >= 0 && tmpI < row) && (tmpJ >= 0 && tmpJ < col) && (board[tmpI][tmpJ] == word[l] && !flag[tmpI][tmpJ]) {
				flag[tmpI][tmpJ] = true
				f = aa(tmpI, tmpJ, l+1)
				flag[tmpI][tmpJ] = false
				if f {
					return true
				}
			}
		}
		return f
	}
	f := false
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				flag[i][j] = true
				f = aa(i, j, 1)
				flag[i][j] = false
				if f {
					return f
				}
			}
		}
	}
	return f
}

// 消耗内存最小
func exist44(board [][]byte, word string) bool {
	m, n, l := len(board), len(board[0]), len(word)
	ans := false
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		if k == l {
			ans = true
			return
		}
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
			return
		}
		temp := board[i][j]
		board[i][j] = byte('-')
		dfs(i+1, j, k+1)
		dfs(i-1, j, k+1)
		dfs(i, j+1, k+1)
		dfs(i, j-1, k+1)
		board[i][j] = temp
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				dfs(i, j, 0)
			}
		}
	}
	return ans
}
