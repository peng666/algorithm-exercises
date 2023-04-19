//package main
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func solution(headA, headB *ListNode) *ListNode {
//	p1, p2 := headA, headB
//	i := 0
//	for i < 3 {
//		if p1 == p2 {
//			return p1
//		}
//		p1 = p1.Next
//		p2 = p2.Next
//		if p1 == nil {
//			p1 = headB
//			i++
//		}
//		if p2 == nil {
//			p2 = headA
//			i++
//		}
//	}
//	return nil
//}
