package main

import (
	"fmt"
	"math/rand"
)

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 快速排序
func quickSort(nums []int, left, right int) []int {
	if left < right {
		pivot := partition(nums, left, right)
		quickSort(nums, left, pivot)
		quickSort(nums, pivot+1, right)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	idx := rand.Intn(right-left) + left
	pivot := nums[idx]
	nums[idx], nums[left] = nums[left], nums[idx]
	for left < right {
		for left < right && pivot <= nums[right] {
			right -= 1
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left += 1
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

func main() {
	nums := []int{110, 100, 0}
	sortArray(nums)
	fmt.Println(nums)
}
