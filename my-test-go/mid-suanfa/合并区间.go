package mid_suanfa

import "sort"

func merge(intervals [][]int) [][]int {
	data := make([][]int, 0, len(intervals))

	keyMap := make([]bool, len(intervals))
	left := 0
	for left < len(intervals) {
		if keyMap[left] {
			left++
			continue
		}
		num1 := intervals[left]
		res := make([]int, 2)
		res[0] = num1[0]
		res[1] = num1[1]
		keyMap[left] = true
		for right := left + 1; right < len(intervals); right++ {
			if keyMap[right] {
				continue
			}
			num2 := intervals[right]
			if res[1] >= num2[0] && res[0] <= num2[1] {
				if res[0] >= num2[0] {
					res[0] = num2[0]
				}
				if res[1] <= num2[1] {
					res[1] = num2[1]
				}
				keyMap[right] = true
			}
		}
		data = append(data, res)
		left++
	}

	return data
}

func merge44(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	index := 0
	var ret [][]int
	ret = append(ret, intervals[0])
	for _, interval := range intervals {
		if interval[0] > ret[index][1] {
			index++
			ret = append(ret, interval)
		} else {
			ret[index][1] = max(interval[1], ret[index][1])
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
