package main

import "fmt"

func sortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(nums, start, mid)
	mergeSort(nums, mid+1, end)

	if nums[mid] <= nums[mid+1] {
		return
	}

	temp := make([]int, end-start+1)
	i, j := start, mid+1
	for k, _ := range temp {
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
	for k, num := range temp {
		nums[start+k] = num
	}
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	res := sortArray(nums)
	fmt.Println(res)
}
