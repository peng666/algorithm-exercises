//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
///* 重排链表结合了几道小题：
//*	快慢指针找中点；反转链表；合并链表
//*/
//func reorderList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil || head.Next.Next == nil {
//		return head
//	}
//
//	// 快慢指针找中点
//	fast, slow := head, head
//	for fast != nil && fast.Next != nil {
//		fast = fast.Next.Next
//		slow = slow.Next
//	}
//	head2 := reverse(slow)
//
//	dummy := new(ListNode)
//	cur := dummy
//	flag := true
//	for head2 != nil && head != nil {	// 两边都要判断
//		if flag {
//			cur.Next = head
//			head = head.Next
//		} else {
//			cur.Next = head2
//			head2 = head2.Next
//		}
//		cur = cur.Next
//		flag = !flag	// 要记得这些点
//	}
//	return dummy.Next
//}
//
//func reverse(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	var newHead *ListNode
//	for head != nil {
//		tmp := head.Next
//		head.Next = newHead
//		newHead = head
//		head = tmp
//	}
//	return newHead
//}
