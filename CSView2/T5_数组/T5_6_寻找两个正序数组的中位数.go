package main

/*
// 中位数，就是找总数一半位置上的数
*/

func Solution(nums1, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	length := m + n
	x1, x2 := 0, 0                   // 中位数指针
	l1, l2 := 0, 0                   // 两个数组指针，谁小谁先走
	for i := 0; i <= length/2; i++ { // 中位数，就是找总数一半位置上的数
		x1 = x2
		if l1 < m && (l2 >= n || nums1[l1] < nums2[l2]) {
			x2 = nums1[l1]
			l1++
		} else {
			x2 = nums2[l2]
			l2++
		}
	}
	if length%2 == 1 {
		return float64(x2)
	}
	return float64(x1+x2) / 2
}
