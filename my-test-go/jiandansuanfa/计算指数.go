package main

/*
给定整数 n ，返回 所有小于非负整数 n 的质数的数量 。
*/

// 超时警告 ,枚举
func isPrime32(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes22(n int) (cnt int) {
	for i := 2; i < n; i++ {
		if isPrime32(i) {
			cnt++
		}
	}
	return
}

// 埃氏筛
func countPrimes(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}

// 线性筛
func countPrimes43(n int) int {
	primes := []int{}
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if i*p >= n {
				break
			}
			isPrime[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}

var (
	mask []bool
	dp   []int
)

// 提前算好结果
func countPrimes33(n int) int {
	return dp[n]

}

func init() {
	count := 5*1000000 + 1
	mask = make([]bool, count)
	dp = make([]int, count)
	dp[0] = 0
	mask[0], mask[1] = true, true
	i := 2
	for i < count {
		for i < count && mask[i] {
			i++
		}
		j := i + i
		for j < count {
			mask[j] = true
			j += i
		}
		i++
	}
	for i := 1; i < count; i++ {
		if !mask[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
}
