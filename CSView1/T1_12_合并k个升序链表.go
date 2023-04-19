//package CSView1
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
///*
//	归并排序：
//	其实就是几个小题合成的：归并排序+合并两个有序链表
// */
//func mergeKLists(lists []*ListNode) *ListNode {
//	length := len(lists)
//	if length == 0 {
//		return nil
//	}
//	if length == 1 {
//		return lists[0]
//	}
//
//	left := mergeKLists(lists[0 : length/2])
//	right := mergeKLists(lists[length/2:])
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
//
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
//	if head1 == nil {
//		cur.Next = head2
//	}
//	if head2 == nil {
//		cur.Next = head1
//	}
//	return dummy.Next
//}
