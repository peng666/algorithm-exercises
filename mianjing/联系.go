package main

import "fmt"

func heapSort(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(&nums, i, len(nums)-1)
	}

	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(&nums, 0, i-1)
	}
	return nums
}

func heapify(nums *[]int, root, end int) {
	for {
		child := root*2 + 1
		if child > end {
			return
		}
		if child < end && (*nums)[child] < (*nums)[child+1] {
			child++
		}
		if (*nums)[root] > (*nums)[child] {
			return
		}
		(*nums)[root], (*nums)[child] = (*nums)[child], (*nums)[root]
		root = child
	}
}
func main() {
	nums := []int{2, 4, 2, 33, 53, 23, 11, 2, 9}
	res := heapSort(nums)
	fmt.Println(res)
}
