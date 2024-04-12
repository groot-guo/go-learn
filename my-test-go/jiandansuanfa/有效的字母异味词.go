package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	keyMap := make(map[string]int)
	for index := range s {
		keyMap[string(s[index])] += 1
	}

	for index := range t {
		value, ok := keyMap[string(t[index])]
		if !ok {
			return false
		}
		value--
		if value < 0 {
			return false
		}
		keyMap[string(t[index])] = value
	}
	return true
}

// 运行速度最快
func isAnagram2(s string, t string) bool {
	hash := make([]int, 26)
	for _, i := range s {
		hash[i-'a']++
	}
	for _, i := range t {
		hash[i-'a']--
	}
	for i := 0; i < len(hash); i++ {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}

// 消耗内存最小
func isAnagram3(s string, t string) bool {

	c1 := [26]int{}
	c2 := [26]int{}
	for _, v := range s {
		c1[v-'a']++
	}
	for _, v := range t {
		c2[v-'a']++
	}
	return c1 == c2
}
