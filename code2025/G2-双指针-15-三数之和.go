package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针
		left, right := i+1, len(nums)-1
		for left < right {
			ans := nums[i] + nums[left] + nums[right]
			if ans == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				left += 1
				right -= 1
			} else if ans < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return res
}

func main() {
	nums := []int{-2, 5, 9, 1, 4, 1, -3, 2, -2}
	res := threeSum(nums)
	fmt.Println(res)
}
