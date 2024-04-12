package main

import (
	"strings"
	"unicode"
)

// 1. 两边对比，复杂度 2n
// 2. 双指针解决，左右两边同时进行，最优解
// 3.

func isPalindrome(s string) bool {
	//tmp:=strings.ToLower(s)
	//return helper(tmp)
	if s == "" {
		return true
	}
	maxLen := len(s)
	res := ""
	for index := range s {
		if unicode.IsLetter(rune(s[index])) {
			res += strings.ToLower(string(s[index]))
		}
		if unicode.IsNumber(rune(s[index])) {
			res += strings.ToLower(string(s[index]))
		}
	}

	res2 := ""
	for index := range s {
		index += 1
		if unicode.IsLetter(rune(s[maxLen-index])) {
			res2 += strings.ToLower(string(s[maxLen-index]))
		}
		if unicode.IsNumber(rune(s[maxLen-index])) {
			res2 += strings.ToLower(string(s[maxLen-index]))
		}
	}
	if res == res2 {
		return true
	}

	return false
}

func helper(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if !isValid(s, lo) {
			lo++
			continue
		}
		if !isValid(s, hi) {
			hi--
			continue
		}
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--

	}
	return true
}

func isValid(s string, idx int) bool {
	if s[idx] >= 'A' && s[idx] <= 'Z' {
		return true
	}
	if s[idx] >= 'a' && s[idx] <= 'z' {
		return true
	}
	if s[idx] >= '0' && s[idx] <= '9' {
		return true
	}
	return false
}
