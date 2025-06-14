package main

import (
	"fmt"
)

func minSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1
	sum := 0
	left := 0
	for i, num := range nums {
		sum += num
		for sum >= target {
			length := i - left + 1
			if length < res {
				res = length
			}
			sum -= nums[left]
			left += 1
		}
	}
	if res > len(nums) {
		return 0
	}
	return res
}

func main() {
	res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	fmt.Print(res)
}
