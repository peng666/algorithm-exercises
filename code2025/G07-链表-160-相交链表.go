package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	isEndA, isEndB := false, false
	pa, pb := headA, headB
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
		if pa == nil && !isEndA {
			pa = headB
			isEndA = true
		}
		if pb == nil && !isEndB {
			pb = headA
			isEndB = true
		}
	}
	return nil
}
