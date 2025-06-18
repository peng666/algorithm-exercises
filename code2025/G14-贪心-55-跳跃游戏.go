package main

func canJump(nums []int) bool {
	cover := 0 // 记录最远可达下标
	for i, num := range nums {
		if i > cover {
			return false
		}
		if i+num > cover {
			cover = i + num
		}
		if cover >= len(nums) { // 可以提前退出
			return true
		}
	}
	return true
}
