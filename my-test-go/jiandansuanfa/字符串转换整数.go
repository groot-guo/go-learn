package main

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

// 错误
func myAtoi(s string) int {
	data := 0

	index := 0
	need := 0
	for index < len(s) {
		if string(s[index]) == "." || string(s[index]) >= "a" || string(s[index]) >= "A" {
			break
		}
		if string(s[index]) == " " {
			index++
			continue
		}

		if need != 0 && (string(s[index]) == "-" || string(s[index]) == "+") {
			return 0
		}
		if string(s[index]) == "+" {
			need = 1
		} else if string(s[index]) == "-" {
			need = -1
		} else if string(s[index]) > "0" {
			num, _ := strconv.Atoi(string(s[index]))
			data = data*10 + num
		}
		index++
	}

	if need < 0 {
		data = need * data
	}

	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1
	if data < MinInt32 {
		return MinInt32
	}
	if data > MaxInt32 {
		return MaxInt32
	}
	return data
}

func myAtoi2(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s)
	//丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}
	//标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')  //字节 byte '0' == 48
		if sign*abs < math.MinInt32 { //整数超过 32 位有符号整数范围
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}

func myAtoi3(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	upp := true
	if s[0] == '-' {
		upp = false
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	res := []byte{}
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			res = append(res, s[i])
		} else {
			break
		}
	}
	if len(res) == 0 {
		return 0
	}
	s = string(res)
	if upp {
		x, _ := strconv.Atoi(s)
		return min(x, math.MaxInt32)
	} else {
		x, _ := strconv.Atoi(s)
		return -min(x, math.MaxInt32+1)
	}

}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 最小用时
func myAtoi4(s string) int {
	if len(s) == 0 {
		return 0
	}
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1

	var retData int
	label := true
	i := 0
	for ; i < len(s) && s[i] == ' '; i++ {

	}
	if i >= len(s) {
		return 0
	}
	if s[i] == '-' {
		label = false
	}
	if s[i] == '-' || s[i] == '+' {
		i++
	}
	if i >= len(s) {
		return 0
	}
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {

		retData = retData*10 + int(s[i]) - 48 //字节 byte '0' == 48
		if label && retData > MaxInt32 {
			return MaxInt32
		} else if !label && retData-1 > MinInt32 {
			return MinInt32
		}

	}
	if !label {
		return -retData
	}
	return retData
}
