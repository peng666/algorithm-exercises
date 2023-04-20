//package CSView1
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
///*
// *	链表排序：归并排序
// *	分治法：用快慢指针把链表拆分，对两个链表进行独立排序，再合并链表
//	其实就是合并了2个小题：快慢指针找链表中点；链表排序+合并链表
// */
//
//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	fast, slow := head, head
//	pre := new(ListNode)
//	for fast != nil && fast.Next != nil {
//		pre = slow
//		fast = fast.Next.Next
//		slow = slow.Next
//	}
//	pre.Next = nil
//	left := sortList(head)
//	right := sortList(slow)
//
//	return merge(left, right)
//}
//
//func merge(head1, head2 *ListNode) *ListNode {
//	if head1 == nil {
//		return head2
//	}
//	if head2 == nil {
//		return head1
//	}
//	dummy := new(ListNode)
//	cur := dummy
//	for head1 != nil && head2 != nil {
//		if head1.Val <= head2.Val {
//			cur.Next = head1
//			head1 = head1.Next
//		} else {
//			cur.Next = head2
//			head2 = head2.Next
//		}
//		cur = cur.Next
//	}
//	if head1 != nil {
//		cur.Next = head1
//	} else if head2 != nil {
//		cur.Next = head2
//	}
//	return dummy.Next
//}
