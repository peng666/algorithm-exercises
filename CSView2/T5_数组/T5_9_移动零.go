//package main
//
//func moveZeros(nums []int) {
//	n := len(nums)
//	if n <= 1 {
//		return
//	}
//	fast, slow := 0, 0
//	for fast != n {
//		for fast != n-1 && nums[fast] == 0 {
//			fast++
//		}
//		nums[slow] = nums[fast]
//		slow++
//		fast++
//	}
//	for slow != n {
//		nums[slow] = 0
//		slow++
//	}
//}
//
//func moveZeros2(nums []int) {
//	n := len(nums)
//	if n <= 1 {
//		return
//	}
//	// 参考快排思想
//	for i, j := 0, 0; i < n; i++ {
//		if nums[i] != 0 {
//			nums[i], nums[j] = nums[j], nums[i]
//			j++
//		}
//	}
//}
