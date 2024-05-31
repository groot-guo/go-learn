package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	index := 0
	s := ""
	minLen := len(strs[0])

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	for index < minLen {
		curStr := strs[0][index]
		for i := 1; i < len(strs); i++ {
			if strs[i][index] != curStr {
				return s
			}
		}
		s += string(curStr)
		index++
	}

	return s
}

func main() {
	s := "1"
	y := "1"
	fmt.Println(s[0] == y[0])
	fmt.Println(s + string(y[0]))

	data := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(data))
}

// 消耗最小的内存
func longestCommonPrefix2(strs []string) string {
	minIndex := 0
	minLength := len(strs[0])
	for i := range strs {
		if len(strs[i]) < minLength {
			minLength = len(strs[i])
			minIndex = i
		}
	}
	for i := range strs[minIndex] {
		for j := range strs {
			if strs[minIndex][i] != strs[j][i] {
				// if i == 0 {
				// 	return ""
				// } else {
				return strs[minIndex][:i]
				// }
			}
		}
	}
	return strs[minIndex]
}
