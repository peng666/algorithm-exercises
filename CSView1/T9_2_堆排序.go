package main

import (
	"fmt"
)

func sortArray(nums []int) []int {
	// quickSort(nums, 0, len(nums)-1)
	heapSort(nums)
	return nums
}

// 堆排序
func heapSort(nums []int) []int {
	// 1. 建堆。大顶堆
	for i := len(nums)/2 - 1; i >= 0; i -= 1 {
		heapify(nums, i, len(nums)-1)
	}

	// 2. 调整堆
	for i := len(nums) - 1; i >= 0; i -= 1 {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i-1) // 注意。上一行已经把下表为i的位置排正确了，已经不是end了（否则它又会排上来）
	}
	return nums
}

func heapify(nums []int, root, end int) {
	if root < 0 || root > end {
		return
	}

	child := root*2 + 1 // 注意。左孩子要加1，右孩子是加2了
	if child > end {
		return
	}

	if child+1 <= end && nums[child] < nums[child+1] {
		child += 1
	}
	if nums[root] > nums[child] {
		return
	}

	nums[root], nums[child] = nums[child], nums[root]
	heapify(nums, child, end)
}

func main() {
	nums := []int{2, 4, 2, 33, 53, 23, 11, 2, 9}
	res := sortArray(nums)
	fmt.Println(res)
}
