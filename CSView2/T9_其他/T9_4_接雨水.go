package main

/*
题目：
	数组的每个元素是该柱子的高度
思路：
	双指针。维护两边的最大高度，每次移动最大高度矮的那边（因为木桶效应）
*/

func Solution(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	res := 0
	l, r := 0, n-1
	maxL, maxR := 0, 0
	for l < r {
		// 维护左右的最大木板
		if nums[l] > maxL {
			maxL = nums[l]
		}
		if nums[r] > maxR {
			maxR = nums[r]
		}
		if maxL < maxR { // 计算max小的那边，因为短木板决定水的高度
			res += maxL - nums[l]
			l++
		} else {
			res += maxR - nums[r]
			r--
		}
	}
	return res
}
