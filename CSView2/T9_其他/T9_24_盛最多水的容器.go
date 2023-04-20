package main

/*
题目：
	数组存有木板的长度。求盛水最多的两块木板所装水的体积
解题：
	双指针+贪心。前后指针，每次移动短木板那边的指针（保留长木板，贪心思想）

*/

func Solution(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	l, r := 0, n-1
	res := 0
	for l < r {
		tmp := 0
		if height[l] > height[r] {
			tmp = (r - l) * height[r]
			r--
		} else {
			tmp = (r - l) * height[l]
			l++
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}
