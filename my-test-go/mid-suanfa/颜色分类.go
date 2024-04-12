package mid_suanfa

// o n 单指针
func swapColors(colors []int, target int) (countTarget int) {
	for i, c := range colors {
		if c == target {
			colors[i], colors[countTarget] = colors[countTarget], colors[i]
			countTarget++
		}
	}
	return
}

func sortColors(nums []int) {
	count0 := swapColors(nums, 0) // 把 0 排到前面
	swapColors(nums[count0:], 1)  // nums[:count0] 全部是 0 了，对剩下的 nums[count0:] 把 1 排到前面
}

// 双指针
func sortColors4(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
