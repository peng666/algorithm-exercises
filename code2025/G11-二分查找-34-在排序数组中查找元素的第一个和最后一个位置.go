package main

import "fmt"

/**
 * 二分查找
 * 34. 在排序数组中查找元素的第一个和最后一个位置
 */
func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			leftBound := searchLeftBound(nums[left:mid+1], target) + left
			rightBound := searchLeftBound(nums[mid:right], target+1) + mid - 1
			return []int{leftBound, rightBound}
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}

func searchLeftBound(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 2, 4, 4, 5, 5, 5, 5, 5, 5, 5, 6, 8, 8, 9}
	result := searchRange(nums, 2)
	fmt.Println(result)
}
