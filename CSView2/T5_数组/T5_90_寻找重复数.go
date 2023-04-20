package main

/*
这道题就是：想象成环形链表，找环的入口
想象快慢指针走的次数相等（想想）
*/

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast { // 找环的的部分，进去环里面
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for fast != slow { // 在环里面相遇，就是重复元素(环的入口）
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}
