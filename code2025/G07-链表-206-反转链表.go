package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2, Next: node1}
	node3 := &ListNode{Val: 3, Next: node2}
	res := reverseList(node3)
	for res != nil {
		fmt.Print(res.Val)
		fmt.Print(" ")
		res = res.Next
	}
}
