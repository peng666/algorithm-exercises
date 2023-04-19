//package main
//
//import "fmt"
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	cur := head
//	for cur.Next != nil {
//		next := cur.Next
//		if next.Val == cur.Val {
//			cur.Next = cur.Next.Next
//		} else { // 这个很重要。如果修改过指针，就不用移动指针啦
//			cur = cur.Next
//		}
//	}
//	return head
//}
//
//func main() {
//	node1 := &ListNode{Val: 1}
//	node2 := &ListNode{Val: 1}
//	node3 := &ListNode{Val: 2}
//	node4 := &ListNode{Val: 2}
//	node1.Next = node2
//	node2.Next = node3
//	node3.Next = node4
//	res := deleteDuplicates(node1)
//	for res != nil {
//		fmt.Printf("%d\t", res.Val)
//		res = res.Next
//	}
//}
