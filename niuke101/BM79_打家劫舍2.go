package main

// 环形的数组。分两种初始状态情况：偷第一家、不偷第一家

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	res := 0
	dp := make([]int, len(nums)+1)
	// 偷第一家，不偷最后一家
	dp[1] = nums[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	res = dp[len(nums)-1]

	// 不偷第一家，偷最后一家
	dp2 := make([]int, len(nums)+1)
	for i := 2; i <= len(nums); i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i-1])
	}
	res = max(res, dp2[len(nums)])

	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
