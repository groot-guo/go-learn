package simplyoffer

/*
	连续子数组的最大值
*/
// 动态规划方案
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 连续的累加，确保每次相交都是最大数，否则重新开始计算
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

/*
礼物的最大值
*/
// 判断思路，就是当前遍历的位置与左邻 和上方相加进行判断，最终右下角算出最大值
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		if i > 0 {
			grid[i][0] += grid[i-1][0]
		} else {
			grid[i][0] = grid[i][0]
		}
		for j := 1; j < len(grid[0]); j++ {
			if i > 0 && grid[i][j]+grid[i][j-1] < grid[i][j]+grid[i-1][j] {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func maxValue2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	data := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		data[0] += grid[i][0]
		for j := 1; j < len(grid[0]); j++ {
			data[j] = max(data[j], data[j-1]) + grid[i][j]
		}
	}

	return data[len(data)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
