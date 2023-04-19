package main

import "sort"

// 注意先排序+去重，使用双指针

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n <= 2 {
		return nil
	}

	sort.Ints(nums) // 排序
	res := [][]int{}

	for i := 0; i < n-1; i++ {
		if nums[i] > 0 { // 提前结束
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}

		for j, k := i+1, n-1; j < k; {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] { // 去重
					j++
				}
				j++
				k--
			} else if tmp < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
