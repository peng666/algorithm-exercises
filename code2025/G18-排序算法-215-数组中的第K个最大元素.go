package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	left := 0
	right := n - 1
	target := n - k
	for {
		pivot := partition(nums, left, right)
		if pivot == target {
			return nums[pivot]
		} else if pivot > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
}

func partition(nums []int, left, right int) int {
	randIdx := left + rand.Intn(right-left+1)
	nums[left], nums[randIdx] = nums[randIdx], nums[left]
	pivot := nums[left]
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
	nums[left] = pivot
	return left
}

func main() {
	res := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	fmt.Println(res)
}
