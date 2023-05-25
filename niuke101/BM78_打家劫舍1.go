package main

/*
动态规划：
1 初始状态。肯定偷取
2 状态转移。要么偷了前面这家，不偷当前这家；要么偷前面的前面那家，偷这家
*/

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	// 初始状态
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	// 状态转移
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
