package main

func search(nums []int, target int) int {
	left, right := 0, len(nums) // 记住二分：左闭右开
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		// 肯定有一边有序，找出来，再判断target，二分到一边去
		if nums[left] < nums[mid] { // 这个情况，左边肯定有序
			if nums[left] <= target && target < nums[mid] {
				right = mid // 记住二分：左闭右开
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right-1] { // 记住二分：左闭右开
				left = mid + 1
			} else {
				right = mid // 记住二分：左闭右开
			}
		}
	}
	return -1
}
