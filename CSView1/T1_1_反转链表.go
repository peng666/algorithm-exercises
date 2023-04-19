//package main
//
//import "fmt"
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	var newHead *ListNode // 创建一个新的链表，不断地头插法
//	for head != nil {
//		tmp := head.Next
//		head.Next = newHead
//		newHead = head
//		head = tmp
//	}
//	return newHead
//}
//
//func main() {
//	node1 := &ListNode{Val: 1}
//	node2 := &ListNode{Val: 2}
//	node3 := &ListNode{Val: 3}
//	node4 := &ListNode{Val: 4}
//	node1.Next = node2
//	node2.Next = node3
//	node3.Next = node4
//	res := reverseList(node1)
//	for res != nil {
//		fmt.Printf("%d\t", res.Val)
//		res = res.Next
//	}
//}
