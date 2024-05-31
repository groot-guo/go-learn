package main

func reverseString(s []byte) {
	res := make([]byte, 0, len(s))
	res = append(res, s...)
	left := 0
	right := len(s) - 1
	for left < len(s) {
		s[left] = res[right]
		right--
		left++
	}

}

// 交换
func reverseString2(s []byte) {
	inner(s, 0, len(s)-1)
}

func inner(s []byte, left, right int) {
	if left >= right {
		return
	}
	inner(s, left+1, right-1)
	s[left], s[right] = s[right], s[left]
}
