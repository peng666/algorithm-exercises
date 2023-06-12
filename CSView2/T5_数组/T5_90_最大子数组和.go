package main

/*
最大子数组和
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
1.动态规划(以下方法用了这个，O(n))
2.分支法（一样是O(n)，因为用了递归，空间O(log n)。但是可以求任意区间的结果）
*/

func maxSubArray(nums []int) int {

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

func main() {

}
