package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findListMid(head)
	head2 := mid.Next
	mid.Next = nil

	left := mergeSort(head)
	right := mergeSort(head2)

	result := mergeTwoLists(left, right)
	return result
}
func findListMid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}

	head := node4
	node4.Next = node2
	node2.Next = node1
	node1.Next = node5
	node5.Next = node3
	node3.Next = node6

	printList(head)
	head = mergeSort(head)
	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(fmt.Sprintf("%d ", head.Val))
		head = head.Next
	}
	fmt.Println()
}
