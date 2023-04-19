//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	end := head
//	for i := 0; i < k; i++ {
//		if end == nil {
//			return head
//		}
//		end = end.Next
//	}
//	newHead := reverse(head, end)
//	head.Next = reverseKGroup(end, k)
//	return newHead
//}
//
//func reverse(start, end *ListNode) *ListNode {
//	if start == nil || start.Next == nil {
//		return start
//	}
//	var newHead *ListNode
//	for start != end {
//		tmp := start.Next
//		start.Next = newHead
//		newHead = start
//		start = tmp
//	}
//	return newHead
//}
