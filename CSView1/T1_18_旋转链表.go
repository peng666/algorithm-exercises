//package CSView1
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func rotateRight(head *ListNode, k int) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	length := 1
//	cur := head
//	for cur.Next != nil {
//		cur = cur.Next
//		length++
//	}
//	cur.Next = head
//	k = k % length // 记得取模~~
//	for i := 0; i < length-k; i++ {
//		cur = cur.Next
//	}
//	head = cur.Next
//	cur.Next = nil
//	return head
//}
