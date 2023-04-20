//package CSView1
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
///*
//	输出环形链表起起始节点
//	快慢指针：当第一次相遇，快指针比慢指针多走了一段，刚好是慢指针路径的两倍；于是，此时头节点跟着慢指针走，当它们相与就是环形入口
//*/
//func detectCycle(head *ListNode) *ListNode {
//	fast, slow := head, head
//	for fast != nil && fast.Next != nil {
//		fast = fast.Next.Next
//		slow = slow.Next
//		if fast == slow {
//			for head != slow {
//				head = head.Next
//				slow = slow.Next
//			}
//			return head
//		}
//	}
//	return nil
//}
