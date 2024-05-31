package mid_suanfa

// 二分法
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}

	return nil
}

// 用两个边界方法
func searchRange44(nums []int, target int) []int {
	// 目标值开始位置：为目标值的左侧边界，无此值则返回比它大的数的左侧边界
	start := findLeftBound(nums, target)
	// 如果开始位置越界 或 目标值不存在，直接返回
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	// 目标值结束位置：为目标值+1的左侧边界-1，无此值则返回比它大的数的左侧边界-1
	end := findLeftBound(nums, target+1) - 1
	return []int{start, end}
}

// 寻找左侧边界的二分查找
func findLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // note: [left, right]
	for left <= right {           // note: 因为 right 是闭区间，所以可以取 =
		mid := left + ((right - left) >> 1) // medium = (left + right) / 2 的优化形式，防止溢出！
		if nums[mid] == target {
			right = mid - 1 // note: 收紧右侧边界以锁定左侧边界
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// 返回左侧边界
	return left // note
}

func searchRange55(nums []int, target int) []int {
	ans := []int{}
	//开始位置：查找第一个>=target的数
	//范围[0...n - 1] + [n表示不存在]
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	ans = append(ans, right) // 第一个>=target的数的下标（不存在为n）

	//结束位置：查找最后一个<=target的数
	//范围[-1表示不存在] + [0...n - 1]
	left, right = -1, len(nums)-1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	ans = append(ans, right) // 最后一个<=target的数（不存在为-1）

	//target出现在[ans[0], ans[1]]
	if ans[0] > ans[1] {
		return []int{-1, -1}
	}
	return ans
}
