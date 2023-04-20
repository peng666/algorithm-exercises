//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//// 解题：双指针。两链表长度相等（双指针走一遍就可以找到）；链表长度不同，双指针走两遍肯定能找到
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headA == nil || headB == nil {
//		return nil
//	}
//	ptr1 := headA
//	ptr2 := headB
//	flag := 0 // 用于标记ptr1走过尾结点的次数
//
//	for flag < 2 {
//		if ptr1 == ptr2 {
//			return ptr1
//		}
//		if ptr1.Next == nil {
//			ptr1 = headB
//			flag++
//		} else {
//			ptr1 = ptr1.Next
//		}
//		if ptr2.Next == nil {
//			ptr2 = headA
//		} else {
//			ptr2 = ptr2.Next
//		}
//	}
//	return nil
//}
