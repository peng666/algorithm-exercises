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
//	newHead := reverse(head, end)     // head包含，end不包含
//	head.Next = reverseKGroup(end, k) // 注意这里的参数
//	return newHead
//}
//
//func reverse(head, end *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	var newHead *ListNode // 注意：该函数就是做经典的反转链表，不需变形
//	for head != end {
//		tmp := head.Next
//		head.Next = newHead
//		newHead = head
//		head = tmp
//	}
//	return newHead
//}
