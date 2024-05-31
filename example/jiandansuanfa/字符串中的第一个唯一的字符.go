package main

func firstUniqChar(s string) int {
	keyMap := make(map[string][]int)
	for i := 0; i < len(s); i++ {
		keyMap[string(s[i])] = append(keyMap[string(s[i])], i)
	}

	for _, value := range s {
		if len(keyMap[string(value)]) == 1 {
			return keyMap[string(value)][0]
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	var ru_array [26]int
	for _, ru := range s {
		ru_array[ru-'a'] += 1
	}
	for i, ru := range s {
		if ru_array[ru-'a'] == 1 {
			return i
		}
	}
	return -1
}
