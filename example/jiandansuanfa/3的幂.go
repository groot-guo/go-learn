package main

import "math"

func isPowerOfThree(n int) bool {
	return helperPower(n) > 0
}

func helperPower(n int) int {
	if n%3 != 0 {
		return n % 3
	}
	return helperPower(n / 3)
}

func isPowerOfThree33(n int) bool {
	if n <= 0 {
		return false
	}

	max := int(math.Pow(3, 19))
	return max%n == 0
}

func isPowerOfThree44(n int) bool {
	return n > 0 && 1162261467%n == 0 // 因为存在范围，求最大幂即可
}
