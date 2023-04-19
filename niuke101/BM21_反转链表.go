//package main
//
//import "fmt"
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func ReverseList(pHead *ListNode) *ListNode {
//	if pHead == nil || pHead.Next == nil {
//		return pHead
//	}
//	var newHead *ListNode
//	for pHead != nil {
//		tmp := pHead.Next
//		pHead.Next = newHead
//		newHead = pHead
//		pHead = tmp
//	}
//	return newHead
//}
//
//func main() {
//	node1 := &ListNode{Val: 1}
//	node2 := &ListNode{Val: 2}
//	node3 := &ListNode{Val: 3}
//	node1.Next = node2
//	node2.Next = node3
//	res := ReverseList(node1)
//	for res != nil {
//		fmt.Printf("%d\t", res.Val)
//		res = res.Next
//	}
//}
