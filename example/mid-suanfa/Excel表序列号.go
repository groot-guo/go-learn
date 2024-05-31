package mid_suanfa

func titleToNumber(columnTitle string) int {
	res, n, chek := 0, len(columnTitle)-1, 1
	for n >= 0 {
		res += int(columnTitle[n]-'A'+1) * chek
		chek *= 26
		n--
	}

	return res
}

// 速度最快
func titleToNumber44(columnTitle string) int {
	var cal, base, cur = 0, 1, 0
	for _, ch := range columnTitle {
		cal += base
		cur *= 26
		//
		cur += int(ch - 'A')
		base *= 26
	}
	return cal + cur
}

// 官方解答
func titleToNumber66(columnTitle string) (number int) {
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		number += int(k) * multiple
		multiple *= 26
	}
	return
}
