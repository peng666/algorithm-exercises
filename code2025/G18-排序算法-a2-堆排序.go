package main

import "fmt"

func sortArray(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {
	// 1. 建堆。大顶堆
	for i := len(nums)/2 - 1; i >= 0; i -= 1 {
		heapify(nums, i, len(nums)-1)
	}

	// 2.调整堆。交换堆顶和末尾元素，缩小规模，再调整为大顶堆。
	for i := len(nums) - 1; i >= 0; i -= 1 {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i-1) // 注意。上面已经把i位置排正确了，已经不是end了。（否则它又被排上来）
	}
}

func heapify(nums []int, root, end int) {
	if root < 0 || root > end {
		return
	}
	child := root*2 + 1 // 左孩子
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
	nums := []int{3, 2, 1, 5, 6, 4}
	res := sortArray(nums)
	fmt.Println(res)
}
