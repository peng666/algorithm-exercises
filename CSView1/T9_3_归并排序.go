package main

import "fmt"

func MergeSort(nums []int, start, end int) []int {
	if start >= end {
		return nums
	}
	mid := (start + end) / 2
	MergeSort(nums, start, mid)
	MergeSort(nums, mid+1, end)

	if nums[mid] <= nums[mid+1] {
		return nums
	}
	temp := make([]int, end-start+1)
	i, j := start, mid+1
	for k := 0; k < len(temp); k++ {
		if i > mid {
			temp[k] = nums[j]
			j += 1
		} else if j > end {
			temp[k] = nums[i]
			i += 1
		} else if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i += 1
		} else {
			temp[k] = nums[j]
			j += 1
		}
	}

	for i = 0; i < len(temp); i++ {
		nums[start+i] = temp[i]
	}
	return nums
}

func main() {
	nums := []int{2, 5, 6, 3, 43, 23, 9, 6}
	MergeSort(nums, 0, len(nums)-1) // 不应该传这么多参数进去，应该就只给个数组
	fmt.Println(nums)
}
