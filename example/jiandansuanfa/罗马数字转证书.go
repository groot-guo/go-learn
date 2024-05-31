package main

func romanToInt(s string) int {
	left := 0
	right := 1
	data := 0
	for right < len(s) {
		leftInt := helperInt(string(s[left]))
		rightInt := helperInt(string(s[right]))
		if leftInt < rightInt {
			data += rightInt - leftInt
			left = right + 1
			right = left + 1
		} else {
			data += leftInt
			left++
			right++
		}
	}
	if left != len(s) {
		data += helperInt(string(s[len(s)-1]))
	}
	return data
}

func helperInt(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}
	return 0
}

// 运行最快
func romanToInt22(s string) (ans int) {
	var romanMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := range s { //字符串类型的遍历s[i]表示每一个字符
		value := romanMap[s[i]]
		if i < len(s)-1 && value < romanMap[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}
