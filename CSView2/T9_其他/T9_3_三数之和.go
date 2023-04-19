package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if tmp == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for nums[l] == nums[l+1] {
					l++
				}
			}
		}

	}
}
