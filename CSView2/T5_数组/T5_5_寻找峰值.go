//package main
//
//func Solution(nums []int) int {
//	l, r := 0, len(nums)-1
//	if r == 0 {
//		return r
//	}
//	if nums[l] > nums[l+1] {
//		return l
//	}
//	if nums[r] > nums[r-1] {
//		return r
//	}
//	for l <= r {
//		mid := (l + r) / 2
//		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
//			return mid
//		}
//		if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
//			l = mid
//		} else {
//			r = mid
//		}
//	}
//	return -1
//}
