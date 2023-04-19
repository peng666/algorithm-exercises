//package CSView1
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//// 快慢指针~
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	// 当不知道要不要操作头节点时，最好搞个虚节点
//	dummy := new(ListNode)
//	dummy.Next = head
//	fast, slow := dummy, dummy
//	for i := 0; i < n; i++ {
//		if fast.Next == nil {
//			return head
//		}
//		fast = fast.Next
//	}
//	for fast.Next != nil {
//		fast = fast.Next
//		slow = slow.Next
//	}
//	slow.Next = slow.Next.Next
//	return dummy.Next
//}
