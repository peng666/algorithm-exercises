package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	for pre != nil && pre.Next != nil {
		if pre.Val == pre.Next.Val {
			pre.Next = pre.Next.Next
			continue
		}
		pre = pre.Next
	}
	return head
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 2}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 4}
	node6 := &ListNode{Val: 9}
	node7 := &ListNode{Val: 9}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	pre := node1
	for pre != nil {
		fmt.Printf("%d ", pre.Val)
		pre = pre.Next
	}
	fmt.Println()
	res := deleteDuplicates(node1)
	pre = res
	for pre != nil {
		fmt.Printf("%d ", pre.Val)
		pre = pre.Next
	}
}
