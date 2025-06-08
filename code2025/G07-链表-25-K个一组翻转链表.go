package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	count := 0
	for head != nil {
		head = head.Next
		count += 1
		if count == k {
			for cur.Next != head {
				temp := cur.Next
				cur.Next = temp.Next
				temp.Next = pre.Next
				pre.Next = temp
			}
			pre = cur
			cur = cur.Next
			count = 0
		}
	}
	return dummy.Next
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	printList(node1)
	res := reverseKGroup(node1, 2)
	printList(res)

}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
