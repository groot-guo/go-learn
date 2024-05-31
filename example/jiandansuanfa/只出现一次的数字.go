package main

// 使用map 效率最高，但是内存消耗最大
func singleNumber(nums []int) int {
	keyMap := make(map[int]int)
	for index := range nums {
		keyMap[nums[index]] += 1
	}

	for key, value := range keyMap {
		if value == 1 {
			return key
		}
	}

	return 0
}

// 学习使用位运算解决

func singleNumber2(nums []int) int {
	res := 0
	for index := range nums {
		res ^= nums[index]
	}

	return res
}
