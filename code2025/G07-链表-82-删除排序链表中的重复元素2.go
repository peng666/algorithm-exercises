package main

import "fmt"

/**
 * 82. 删除排序链表中的重复元素 II
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	pre := dummy
	cur := dummy.Next

	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}
		if pre.Next != cur {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{2, nil}
	node4 := &ListNode{3, nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	tmp := node1
	for tmp != nil {
		fmt.Print(tmp.Val, " ")
		tmp = tmp.Next
	}
	fmt.Println()
	deleteDuplicates(node1)

	tmp = node1
	for tmp != nil {
		fmt.Print(tmp.Val, " ")
		tmp = tmp.Next
	}
}
