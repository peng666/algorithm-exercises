package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	tail, pre := head, dummy
	for tail != nil && n > 0 {
		tail = tail.Next
		n -= 1
	}
	for tail != nil {
		tail = tail.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}
