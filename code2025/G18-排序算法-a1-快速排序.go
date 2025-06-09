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
	randIdx := left + rand.Intn(right-left+1)
	nums[left], nums[randIdx] = nums[randIdx], nums[left]
	pivot := nums[left] // 基准元素拿出来
	for left < right {
		for left < right && nums[right] >= pivot {
			right -= 1
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left += 1
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot // 基准元素放在正确位置上
	return left
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	res := sortArray(nums)
	fmt.Println(res)
}
