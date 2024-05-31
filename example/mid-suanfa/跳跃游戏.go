package mid_suanfa

func canJump(nums []int) bool {
	var p int
	for i := range nums {
		if i > p {
			return false
		}
		p = maxInt(p, i+nums[i])
	}
	return true
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
