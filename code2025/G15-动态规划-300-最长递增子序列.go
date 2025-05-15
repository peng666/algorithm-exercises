package main

import "fmt"

func Solution(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	res := dp[0]
	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func Solution2(nums []int) int {
	tails := make([]int, len(nums))
	res := 0
	for _, num := range nums {
		left, right := 0, res
		for left < right {
			mid := (left + right) / 2
			if tails[mid] >= num { // 这里想清楚边界。相等的值不要，所以放在外面
				right = mid
			} else {
				left = mid + 1
			}
		}
		tails[left] = num
		if right == res {
			res += 1
		}
	}
	return res
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	//nums := []int{10, 9, 2, 5, 3, 7, 110, 38}
	nums := []int{9, 9, 9}
	res := Solution2(nums)
	fmt.Println(res)
}
