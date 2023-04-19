//package main
//
//func Solution(nums []int, k int) int {
//	k = len(nums) - k // 第k大的数的下标
//	l, r := 0, len(nums)-1
//	idx := 0
//	for {
//		idx = parition(nums, l, r)
//		if k == idx {
//			return nums[idx]
//		}
//		if k < idx {
//			r = idx - 1
//		} else {
//			l = idx + 1
//		}
//	}
//	return -1
//}
//
//func parition(nums []int, start, end int) int {
//	if start > end {
//		return -1
//	}
//	l, r := start, end
//	pivot := nums[l]
//	for l < r {
//		for l < r && nums[r] >= pivot {
//			r--
//		}
//		nums[l] = nums[r]
//		for l < r && nums[l] <= pivot {
//			l++
//		}
//		nums[r] = nums[l]
//	}
//	nums[l] = pivot
//	return l
//}
