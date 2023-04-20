//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func mergeTwoLists(list1, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	} else if list2 == nil {
//		return list1
//	}
//	dummy := new(ListNode)
//	cur := dummy
//	for list1 != nil && list2 != nil {
//		if list1.Val <= list2.Val {
//			cur.Next = list1
//			list1 = list1.Next
//		} else {
//			cur.Next = list2
//			list2 = list2.Next
//		}
//		cur = cur.Next
//	}
//	if list1 == nil {
//		cur.Next = list2
//	}
//	if list2 == nil {
//		cur.Next = list1
//	}
//	return dummy.Next
//}
