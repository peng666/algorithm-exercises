package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	// 从后往前扫
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	if n > 0 {
		nums1 = append(nums1[:0], nums2[:n]...) // 这里比较有技巧，想清楚底层原理
	}
}

func main() {
	nums1 := []int{1, 4, 7, 0, 0, 0, 0}
	nums2 := []int{3, 5, 9, 12}
	merge(nums1, 3, nums2, 4)
	fmt.Println(nums1)
}
