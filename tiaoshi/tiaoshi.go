package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	nums := []int{2, 4, 3, 0, -2, 8, 23, 8, 433, 22, 3, 4}
	res := findKthLargest(nums, 5)
	fmt.Println(res)
}
func findKthLargest(nums []int, k int) int {
	//k = len(nums) - k
	l, r := 0, len(nums)-1
	for {
		idx := partition(nums, l, r)
		if idx == len(nums)-k {
			return nums[idx]
		} else if idx > k-1 {
			r = idx - 1
		} else {
			l = idx + 1
		}
	}

}

func partition(nums []int, l, r int) int {
	if l > r {
		return -1
	}
	pivot := nums[l]
	for l < r {
		for l < r && pivot < nums[r] {
			r--
		}
		nums[l] = nums[r]
		for l < r && pivot >= nums[l] {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	return l
}
