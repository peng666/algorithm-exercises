package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	dummy := &ListNode{}
	cur := dummy
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 != nil {
		cur.Next = head1
	}
	if head2 != nil {
		cur.Next = head2
	}
	return dummy.Next
}
