package main

func maxSubArray(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	maxInt := dp[0]
	for i := 1; i < length; i++ {
		dp[i] = max(dp[i-1], 0) + nums[i]
		maxInt = max(maxInt, dp[i])
	}

	return maxInt
}

func maxSubArray222(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 分治
func maxSubArray33(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

type Status struct {
	lSum, rSum, mSum, iSum int
}
